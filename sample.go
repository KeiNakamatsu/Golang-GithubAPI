package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"syscall"

	"github.com/google/go-github/github"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Print("GitHub Username: ")
	username, _ := r.ReadString('\n')

	fmt.Print("GitHub Password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	password := string(bytePassword)

	tp := github.BasicAuthTransport{
		Username: username,
		Password: password,
	}

	client := github.NewClient(tp.Client())
	ctx := context.Background()

	pr, res, err2 := client.PullRequests.Get(ctx, "sri-nanki", "ism_Server", 1)
	if err2 != nil {
		fmt.Printf("\n\n%v\n", err2)
		return
	}
	if res.NextPage > 100 {
		return
	}
	fmt.Printf("\n\n")
	fmt.Printf(pr.GetTitle() + "\n")
	fmt.Printf(pr.GetState() + "\n")

}
