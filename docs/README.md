# faq-service
The service is responsible for store frequently asked questions and answers. Provides simple CRUD API.

This project implements [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) in order to assure both high elasticity and reusability of the code.
#### Layers

- `internal/api` - can import any other layers.
- `internal/infrastructure` - can import everything except `api`.
- `internal/useCases` - should not import anything else than `internal/app` or standard library (or `utils`)
- `internal/app` - should not import anything else than the standard library or `utils`

#### Middleware

In order to reuse common code such as authorization and graceful shutdown this project utilizes also middleware
pattern. Handlers are wrapped in pipelines where each step wraps and calls the next one.

## API usage:

**Create FAQ object:**
```text
POST /faq
    {
        "question":"What is your name?",
        "answer": "My name is Bot. Chat Bot :)"
    }
```

**Update FAQ object by ID:**
```text
PATCH /faq/:id
    {
        "question":"How are you, mr Bot?",
        "answer": "I'm fine."
    }
```

**Get FAQ object by ID:**
```text
GET /faq/:id
```

**Delete FAQ object by ID:**
```text
DELETE /faq/:id
```