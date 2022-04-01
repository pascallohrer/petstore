# Backend Engineer - Take Home Assignment

<center><img src="https://www.valutachange.de/wp-content/uploads/2021/09/val-benefits_b1-e1632074593727-1024x781.png" width="200"/></center>

Engineers at ValutaChange build highly available distributed systems in a microservices environment. Our take home test is designed to evaluate real world activities that are involved with this role. We recognise that this may not be as mentally challenging and may take longer to implement than some algorithmic tests that are often seen in interview exercises. Our approach however helps ensure that you will be working with a team of engineers with the necessary practical skills for the role (as well as a diverse range of technical wizardry).

# Instructions

The goal of this exercise is to build a small, self-contained web API in Go, fulfilling a part of the included OpenAPI 3.0 pet store specification provided by [swagger.io](https://swagger.io). There is no need to interface with a real data provider, you're free to maintain data in memory only. Please note that authorisation and authentication are outside the scope of this assignment.

Starting your design from scratch, your submission should:

- be written in Go
- serve the `POST`, `GET {petId}` and `DELETE` operations on the `pet` resource as defined in the `spec/petstore.yaml` document
- perform at least some form of data verification
- follow [SOLID](https://en.wikipedia.org/wiki/SOLID), [DRY](https://en.wikipedia.org/wiki/Don%27t_repeat_yourself) and [KISS](https://en.wikipedia.org/wiki/KISS_principle) principles
- contain tests on a level you'd expect to see on a commercial solution
- run tests and coverage on `make test`, build a docker image on your binary using `make build` and running that binary, ready to accept incoming connections by calling `make run`

# Out of scope

The following topics are not within the scope of this assignment, including one or more of them as part of your submission will not increase your likelyhood of passing:

- API versioning
- Scaling out
- Authorisation & authentication
- Database integration or principles like [ACID](https://en.wikipedia.org/wiki/ACID)
- Securing and hardening your endpoints beyond best practice recommendations

You're free to include your thoughts on how to address one or more of these or any other topics in your submission, especially if this benefits the understanding of your submission, but your contributions will not be evaluated as part of your final solution.

# Submission Guidance

- Make sure to read the given requirements **carefully**
- **Take your time** to design a solution that best meets the given criteria outlined above
- Please see the included `Makefile` for details on how your submission will be handled, our engineers will call `make test` and `make build && make run`, make sure your solution will both be able to run tests and the binary without any additional steps.
- Make sure to include any assumptions made, further documentation and comments needed for understanding your submission
- When ready, upload your submission to a code hosting provider (i.e [Github](https://github.com) or [Gitlab](https://gitlab.com)) and share it with our development team, either by sending a link or inviting `@vctech-recruit`
