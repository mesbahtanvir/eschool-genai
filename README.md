# Learning Track

Learning should be fun and effective, there is opportunity to improve learning experience.

## The Product

### Problem that we are trying to solve

User want to learn about something, we would like to provide tool to help user learn faster and effectively.

## Development

### Local Setup

Requirements

- Docker
- OPEN AI Account

Export OpenAI API Key

```
export OPENAI_API_KEY="YOUR_OPEN_AI_API_KEY"
```

Before executing these make sure docker daemon is up and running.

```
docker compose build
docker compose up -d
```

Visit the locally deployed frontend and play around with it

### Architecture

```
Frontend (reactjs) --> Backend(golang)  --> LLM
                                       |
                                       |
                                        --> MongoDB
```

#### Frontend

Frontend has to be very light weight and clutter free. on left panel it will list down all the current track the user has in progress. once clicked on one of the listed track. It will show last few course content and will have a button named "Advance". Once clicked on Advance button it will show next information.

#### Backend

The backend will have three specific role in the system.

1. Facilitate storing & retriving user data generated in this system.
2. Facilitate retriving user data generated in external sources.
3. Generate Appropiate [prompt](./PROMPT.md) for the LLM

#### LLM

Generator of new content based on existing information about user. At this day, LLM might not be good at doing what we need. But we hope eventually LLM will improve and be better at this.

#### MongoDB

Storing users data, course and other stuff.

Here how it works.

- User want to learn about Linear Algebra, user put `Linear Algebra` as input in the search bar.
- Backend receive the string `Linear Algebra`.
- Backend retrive all the user data related to Linear Algebra
- Backend create a [prompt](./PROMPT.md), example prompt `Here what I know about Linear Algebra, %ALL_THE_DATA_USER_KNOW_ABOUT_LINEAR_ALGEBRA_AS_STRING%. Create a Learning Track Preview where I know to pass Linear Algebra exam. The preview should Contain Title, Subtitle, List of topics the Learning track will cover` and send it to LLM
- Backend Store the response and send it to User.
- User select the learning Track and the Journey Begins
