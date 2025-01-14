# ia-exercise 

## Installation 

```bash
go get -u
go mod tidy
go mod download
```

## Running 

For the project to run properly you need to provide an OpenAI API KEY. The name given to the env var is "OPENAI_API_KEY". 

For running the project use the following comand: 

```bash
go run ./app
```

## Description 
This project pretends to be a chat where given a sentence or a question from a user input, the system respond with a short fun fact from the main topic of the input and using the same slang an language than the one used by de user. 

## Future imporvements 
For now, the project dosn't handel user data, although in the chat input there is parameter for the user id. A future improvement could be the user management, and saving all the topics the user has commented, in order to offer more user appealing fun facts. 


## Aditional information 
As required, the project is containerised. I tried my best in this part but it's not my key point. 

It's not clear for me if I had to do testing of the developed code. I did not do this part, but if needed, I would test service.question.go. I wold do the folowing steps: 

1. Create a file called service.quetions_test.go next to the file I'll test.
2. Mock all repositories, proxies and other services that uses the service we are going to test. In this case, there only needs two proxies. 
3. Write the desired test. One possible test woild be to check if the events from de diferent channels are sent properly. 
4. Run tests. 

For testing other parts of the project, like the openAI proxy, I'd do the same process, changing the file name. 

