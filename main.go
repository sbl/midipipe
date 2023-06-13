package main

import (
	"fmt"
	"os"

	"gitlab.com/gomidi/midi/v2"
	_ "gitlab.com/gomidi/midi/v2/drivers/portmididrv"
	"gitlab.com/gomidi/midi/v2/smf"
)

func main() {
	var midifile *os.File

	if len(os.Args) > 1 {
		// read either from file
		var err error
		midifile, err = os.Open(os.Args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		// or from stdin
		midifile = os.Stdin
		stat, err := midifile.Stat()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if stat.Size() == 0 {
			fmt.Println("stdin seems to be emtpy, quitting.")
			os.Exit(1)
		}
	}

	defer midi.CloseDriver()

	out, err := midi.FindOutPort("IAC Driver Bus 1")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	reader := smf.ReadTracksFrom(midifile)
	if err := reader.Play(out); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
