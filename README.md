# riva_speech
``` go
package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	pb "github.com/your-username/my-riva-pb/pb" // Import your generated Riva STT package
)

// RivaSTT listens for audio on sttGRPC, sends it to Riva, and writes transcriptions to a file.
func RivaSTT(sttGRPC <-chan []byte, outputFile string) {
	// Connect to Riva gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to connect to Riva:", err)
	}
	defer conn.Close()

	client := pb.NewRivaSpeechRecognitionClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Open file for writing transcriptions
	file, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Failed to open file:", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	// Create gRPC streaming request
	stream, err := client.StreamingRecognize(ctx)
	if err != nil {
		log.Fatal("Failed to create gRPC stream:", err)
	}

	// Send initial config
	err = stream.Send(&pb.StreamingRecognizeRequest{
		StreamingRequest: &pb.StreamingRecognizeRequest_Config{
			Config: &pb.RecognitionConfig{
				Encoding:        pb.AudioEncoding_LINEAR_PCM,
				SampleRateHertz: 16000, // Adjust based on your Freeswitch config
				LanguageCode:    "en-US",
			},
		},
	})
	if err != nil {
		log.Fatal("Failed to send config:", err)
	}

	// Goroutine to receive transcription results
	go func() {
		for {
			resp, err := stream.Recv()
			if err != nil {
				log.Println("Error receiving from Riva:", err)
				return
			}
			for _, result := range resp.Results {
				if result.IsFinal {
					transcription := result.Alternatives[0].Transcript
					fmt.Println("Transcription:", transcription)

					// Write transcription to file
					_, err := writer.WriteString(transcription + "\n")
					if err != nil {
						log.Println("Error writing to file:", err)
					}
					writer.Flush()
				}
			}
		}
	}()

	// Listen for audio on the channel and send it to Riva
	for audio := range sttGRPC {
		err := stream.Send(&pb.StreamingRecognizeRequest{
			StreamingRequest: &pb.StreamingRecognizeRequest_AudioContent{
				AudioContent: audio,
			},
		})
		if err != nil {
			log.Println("Error sending audio:", err)
			break
		}
	}

	// Close the stream when done
	stream.CloseSend()
}
```
How to Use This Function

    Start the RivaSTT goroutine:
``` go
sttGRPC := make(chan []byte, 10) // Buffered channel for audio chunks
go RivaSTT(sttGRPC, "transcriptions.txt")
```
Send L16 audio to sttGRPC: Assuming you are reading 4096-byte chunks from FreeSWITCH over WebSocket:
```
for {
    audioChunk := make([]byte, 4096) // Replace with actual received audio data
    // Read from WebSocket and send to channel
    sttGRPC <- audioChunk
}
```


The transcriptions will be saved in transcriptions.txt.
