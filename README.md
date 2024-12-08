### Build

`docker compose build`

### Run Docker Service

`docker compose up -d`

### Problem that we are trying to solve

User want to learn about something

- `Concept`/`Fact`/`Event`: This is self explanatory.
- `Interest`: The `concept`/`fact`/`event` user want to learn. For example, User might want to improve their Linear Algebra understanding, learn basic 3D geometry, WWI or WWII
- `Outcome`: The end goal of user for the Interest. For example, user want to achieve sufficient knowledge on WWI to understand current global politics, Pass GRE/IETLS with certain score.
- `Pre Existing Knowledge`: What user already know about the `Interest`. For example user want to learn Linear Algebra(LA), User already know about basic matrix operation, basic understanding on food and nutrition. The Prexisting knowledge here is basic matrix operation as this is the only pre-existing knowlege that is related to LA.
- `Fractional Information`: A small piece of information that can be meaninfully presented to the user.
- `Knowledge Test`: A `quiz`/`puzzle`/`interactive game` to test user understanding about a `Fractional Information` related to the `Interest` which has already been presented to the user
- `Pre Existing Knowledge Test`: A `quiz`/`puzzle`/`interactive game` to test user understanding about a `Fractional Information` related to the `Interest` which has _not_ yet been presented to the user
- `Optimally`: Optimal strategy are the strategy where user observe low cognitive fatigue, high engagement to achieve their desired outcome in minimal time.

- `Learning Track`: A product to help user to optimally obtain their desired `Outcome` on the `Interest`, considering their `Pre Existing Knowledge`.
