package main

import (
	"net/http"

	"github.com/carloscasalar/traveller-rpg-api/internal/npc"

	"github.com/syumai/workers"
)

func main() {
	http.HandleFunc("/api/npcs/single", npc.SingleHandler)

	workers.Serve(nil) // use http.DefaultServeMux
}
