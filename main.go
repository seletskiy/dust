package main

import (
	"fmt"
	"os"

	"github.com/godbus/dbus"
)

// https://developer.gnome.org/notification-spec/

type Server struct{}

func (server *Server) GetServerInformation() (
	name string,
	vendor string,
	version string,
	spec string,
	err *dbus.Error,
) {
	return "dust", "seletskiy", "0.1", "1.2", nil
}

func (server Server) Notify(
	name string,
	id uint,
	icon string,
	summary string,
	body string,
	actions []string,
	hints map[string]dbus.Variant,
	expire int,
) (uint, *dbus.Error) {
	fmt.Fprintf(os.Stderr, "XXXXXX eavesdrop.go:53 name: %#v\n", name)
	fmt.Fprintf(os.Stderr, "XXXXXX eavesdrop.go:54 id: %#v\n", id)
	fmt.Fprintf(os.Stderr, "XXXXXX eavesdrop.go:55 icon: %#v\n", icon)
	fmt.Fprintf(os.Stderr, "XXXXXX eavesdrop.go:56 summary: %#v\n", summary)
	fmt.Fprintf(os.Stderr, "XXXXXX eavesdrop.go:57 body: %#v\n", body)
	fmt.Fprintf(os.Stderr, "XXXXXX eavesdrop.go:58 actions: %#v\n", actions)
	fmt.Fprintf(os.Stderr, "XXXXXX eavesdrop.go:59 hints: %#v\n", hints)
	fmt.Fprintf(os.Stderr, "XXXXXX eavesdrop.go:60 expire: %#v\n", expire)
	return 0, nil
}

func main() {
	conn, err := dbus.SessionBus()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to session bus:", err)
		os.Exit(1)
	}

	reply, err := conn.RequestName("org.freedesktop.Notifications", dbus.NameFlagDoNotQueue)
	if err != nil {
		panic(err)
	}

	if reply != dbus.RequestNameReplyPrimaryOwner {
		fmt.Fprintln(os.Stderr, "name already taken")
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "XXXXXX eavesdrop.go:193 reply: %#v\n", reply)

	var server Server

	err = conn.Export(&server, "/org/freedesktop/Notifications", "org.freedesktop.Notifications")
	if err != nil {
		panic(err)
	}

	select {}
}
