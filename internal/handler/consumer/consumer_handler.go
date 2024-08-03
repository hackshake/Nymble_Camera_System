package consumer_handler

/*here we register our consumer with the topic
for example:

RegisterHandler(
	--topic,
	--func,
	--message format
)

This will be called from the Start Consumer function

func ProcessConsumerMessgae(message string) {
	 go func(message string) {
	 	request = deserialiseMessage(message)
	 	cameraController.StartCapture(message)
            }(message )
        }
    }()
}
*/
