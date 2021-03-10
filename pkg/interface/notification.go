package interface

import (
	"log"

	"github.com/gen2brain/beeep"
	md "github.com/sonr-io/core/internal/models"
)

const kCompatibleFileTypes = "*.png *.jpg *.jpeg *.mp4 *.avi *.pdf *.doc *.docx *.ttf *.mp3 *.xml *.csv *.key *.ppt *.pptx *.xls *.xlsm *.xlsx *.rtf *.txt"

func PushInvited(inv *md.AuthInvite) {
	err := beeep.Notify("Invited", inv.From.Profile.FirstName+" has sent an invite to share "+inv.Payload.String(), "assets/information.png")
	if err != nil {
		log.Println(err)
	}
}

func BeepCompleted() {
	err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	if err != nil {
		log.Println(err)
	}
}

// ^ Presents a Authentication Dialog for Approval ^ //
func ShowAuthDialog(inv *md.AuthInvite) bool {
	// Set Text
	description := "Do you accept " + inv.From.Profile.FirstName + "'s invite to receive an " + inv.Payload.String()

	// Display
	decision, err := dlgs.Question("Sonr Invitation", description, true)
	if err != nil {
		log.Panicln(err)
	}

	// Return
	return decision
}

// ^ Presents a Error Message ^ //
func ShowErrorDialog(errorMsg *md.ErrorMessage) {
	_, err := dlgs.Error("Error", errorMsg.Message)
	if err != nil {
		panic(err)
	}
}

// ^ Presents a File Input Dialog ^ //
func ShowFileDialog() string {
	// Display
	filename, _, err := dlgs.File("Select a File to Send...", kCompatibleFileTypes, false)
	if err != nil {
		log.Panicln(err)
	}

	// Return
	return filename
}

// ^ Presents a URL Input Dialog ^ //
func ShowURLDialog() string {
	// Display
	url, _, err := dlgs.Entry("URL Link", "Enter a URL Here: ", "")
	if err != nil {
		log.Println(err)
	}

	// Return
	return url
}

// ^ Presents a URL Input Dialog ^ //
func ShowNameDialog() string {
	// Display
	url, _, err := dlgs.Entry("Device Name", "Enter a Device Nickname Here: ", "")
	if err != nil {
		log.Println(err)
	}

	// Return
	return url
}