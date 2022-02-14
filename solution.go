package solution

import "github.com/kyokomi/emoji/v2"

func GetMessage() string {
	stringWithEmoji := emoji.Sprint(":clown_face: your face when...")
	return stringWithEmoji
}
