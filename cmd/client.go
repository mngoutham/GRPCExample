package main

import (
	"context"
	"flag"
	"fmt"
	todo "github.com/mngoutham/golang_basics/build/gen"
	grpc "google.golang.org/grpc"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "missing subcommand: list or add")
		os.Exit(1)
	}

	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not connect to backend: %v\n", err)
		os.Exit(1)
	}
	client := todo.NewTasksClient(conn)

	switch cmd := flag.Arg(0); cmd {
	case "add":
		err = add(context.Background(), client, strings.Join(flag.Args()[1:], " "))
	default:
		err = fmt.Errorf("unknown subcommand %s", cmd)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func add(ctx context.Context, client todo.TasksClient, text string) error {
	_, err := client.Add(ctx, &todo.Text{Text: text})
	if err != nil {
		return fmt.Errorf("could not add task in the backend: %v", err)
	}

	fmt.Println("task added successfully")
	return nil
}
