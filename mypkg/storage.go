package mypkg

import (
  "time"
  "os"
)

mailsFolderName := "mails"

type EmailAddr string
type UserMails map[string][]Message

type MailBox struct {
  // domain string 
  mails UserMails
}


type Message struct{
  To []EmailAddr
  From EmailAddr
  Body string
  Subject string
  Date time.Time
  Deleted bool
}

// Get a user's mail list
func (uM *UserMails) GetMails(user EmailAddr) int {

  return 1
}

// Get a user's mail
func (uM *UserMails) GetMail(user EmailAddr, num int) int {
  return 17
}

// Add a mail to a user mailbox
func (uM *UserMails) AddMail(user EmailAddr, message Message) int {
  return 1
}

func EnsureDirectoryExist(name string) {
  _, err := os.Stat(mailsFolderName)
  if os.IsNotExist(err) {
  }
}
