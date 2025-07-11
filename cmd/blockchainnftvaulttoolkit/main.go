// cmd/blockchainnftvaulttoolkit/main.go
package main

import (
"flag"
"log"
"os"

"blockchainnftvaulttoolkit/internal/blockchainnftvaulttoolkit"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := blockchainnftvaulttoolkit.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
