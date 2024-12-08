### Build

`docker compose build`

### Run Docker Service

`docker compose up -d`

### Problem that we are trying to solve

User want to learn about something, we would like to provide tool to help user learn faster, effectively

### Prompt Engineering

Given The Definition

DEFINITION_START_HERE

- `Concept`/`Fact`/`Event`: Concept: Some abstract idea including math, science. Fact: Something that can be proved based on some concrete event Event: Something that already taken place and historically accurate.
- `Interest`: The `concept`/`fact`/`event` user want to learn. For example, User might want to improve their Linear Algebra understanding, learn basic 3D geometry, World War 1 (WWI) or World Ward 2 (WWII)
- `Outcome`: The end goal of user for the `Interest`. For example, user want to achieve sufficient knowledge on WWI to understand current global politics, Pass GRE/IETLS with certain score.
- `Pre-existing Knowledge`: What user already know about the `Interest`. For example user want to learn Linear Algebra(LA), User already know about basic matrix operation, basic understanding on food and nutrition. The `Pre-existing knowledge` here is basic matrix operation as this is the only information that is related to LA.
- `Fractional Information`: A small piece of information that can be meaninfully presented to the user.
- `Quiz`: Some form of test, where a question will be asked and some option will be provided where only only option answer the question. The Goal of `Quiz` is to understand user's understanding of the Fractional Information/Interest
- `Puzzle`: A form of brain teaser, visual or text based, where user has to use their cognitive ability to solve it.
- `Interactive Game`: A text based game designed to reinformation knowledge or improve cognitive ability.
- `Knowledge Test`: A `Quiz`/`puzzle`/`Interactive Game` to test user understanding about a `Fractional Information` related to the `Interest` which has already been presented to the user
- `Pre-existing Knowledge Test`: A `quiz`/`puzzle`/`interactive game` to test user understanding about a `Fractional Information` related to the `Interest` which has _not_ yet been presented to the user
- `Test`: Either `Knowledge Test` or `Pre-existing Knowledge Test`
- `Optimally`: Optimal strategy are the strategy where user observe low cognitive fatigue, high engagement to achieve their desired outcome in minimal time.
- `Learning Track`: A product to help user to optimally obtain their desired `Outcome` on the `Interest`, considering their `Pre Existing Knowledge`.

DEFINITION_END_HERE

User's Interest

--- ATTACH USER INTEREST HERE ---

User's Pre-existing Knowledge

--- ATTACH PRE EXISTING KNOWLEDGE HERE ---

Given Current Learning Track State with User's Answer

--- ATTACH CURRENT LEARNING TRACK STATE ---

Analyze The given information and determine should user need new Fractional Information or Test.

Now present new `Fractional Information`/`Quiz` to help achieve `Outcome` Considering Pre Existing User knowledge provided above.
