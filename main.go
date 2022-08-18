package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	content, err := os.ReadFile("transcript.json")
	if err != nil {
		log.Fatal(err)
	}

	tr := &Transcript{}
	err = json.Unmarshal(content, tr)
	if err != nil {
		log.Fatal(err)
	}

	transcript := ""

	var lastSpeaker *SpeakerLabel
	itemIdx := 0
	// the assumption is that segments and items are sorted in ascending order by start date  (based on observations not documentation)
	for _, s := range tr.Results.SpeakerLabels.Segments {
		end, err := strconv.ParseFloat(s.EndTime, 32)
		if err != nil {
			log.Fatal(err)
		}

		currentSpeaker := s.SpeakerLabel

		if lastSpeaker == nil || currentSpeaker != *lastSpeaker {
			transcript = fmt.Sprintf("%s\n%s: ", transcript, currentSpeaker)
			lastSpeaker = &currentSpeaker
		}

		for i := itemIdx; i < len(tr.Results.Items); i++ {
			item := tr.Results.Items[i]
			if item.StartTime == nil || item.EndTime == nil {
				continue
			}

			itemStart, err := strconv.ParseFloat(*item.StartTime, 32)
			if err != nil {
				log.Fatal(err)
			}

			// if item is in current segment
			if itemStart < end {
				// take first alternative its the one with highest confidence
				transcript = fmt.Sprintf("%s %s", transcript, item.Alternatives[0].Content)
			} else {
				itemIdx = i
				break
			}
		}
	}

	fmt.Print(transcript)
}
