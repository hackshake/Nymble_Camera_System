# Nymble_Camera_System
A Camera System which captures images concurrently

HLD - diagrams/HLD.png
Sequence Diagram - diagrams/Sequence.png

I didn't have time to create the LLD diagram but I can explain the same in the flow:-
1. Start the app, this does all the necessary dependency injection and start the consumers and starts the server
2. Client calls api "capture" which invokes - http.HandleFunc("/capture", cameraHandler.CaptureImageHandler())
3. cameraHandler.CaptureImageHandler() then calls csh.service.SubmitCaptureRequest(&request) with registered callbacks
4. This then calls Request Manager to Addrequest which publishes the message to the urgency based topic
5. Consumer is listening to this topic and fetches the message and then calls Camera Controller for Capturing the Image which on Capture invokes the callback.

Kafka:-
1. We can have priority based topics and we can increase the consumer count on the same basis
High priority topics will have more consumer counts and so on
2. Also another improvisation can be done while consuming message from kafka. We can add the process messages to channels which can we listened and then those events can be sent to Camera Controller. This can add another layer of decoupling which makes it easier to manage the rate of consumption and processing independently.That can also improve parallelisation and provide back pressure handling.

Suggestion:- Instead of making the client wait we can also send captured data via websockets
