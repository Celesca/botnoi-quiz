# Documentation

## 2. RestAPI for PokeIndex

- API Endpoints is POST `http://localhost:8080/pokemon/:id`

## 3. Line Chatbot with Golang

* ### Installation
  - Create new Line OA
  - Change `channelSecret` and `channelToken` in the code line 16-17.
  - Start ngrok with command `ngrok http 8080`
  - Go to Messaging API settings and replace Webhook URL with ngrok URL
 
* ### Features
  - Text Message -> Typing `"Text"` to chatbot
  - Button Template Message -> Typing `"Button"` to chatbot
  - QuickReply Message -> Typing `"QuickReply"` to chatbot
  - Carousel Template Message -> Typing `"Carousel"` to chatbot
