package main

import (
	"voice-assistant/openai"
	"voice-assistant/speech_to_text"
	"voice-assistant/text_to_speech"
)

func main() {
	// 1. Record user's voice input (or use a pre-recorded file)
	audioFile := "path/to/audio/file.wav"

	// 2. Convert the voice input to text
	userInput := speech_to_text.SpeechToText(audioFile)

	// 3. Send the text input to ChatGPT via OpenAI API
	responseText := openai.ChatGPT(userInput)

	// 4. Convert the text response to speech
	text_to_speech.TextToSpeech(responseText)
}
