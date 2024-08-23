*Question:*

You are tasked with designing a RESTful API for a user management system. The system needs to handle CRUD operations for users and should follow best practices for API design, including adherence to SOLID principles and design patterns.

1.⁠ ⁠*Requirements:*
- The API should support creating, reading, updating, and deleting user profiles.
- Each user profile has an ⁠ id ⁠, ⁠ name ⁠, ⁠ email ⁠, ⁠ password ⁠, and ⁠ role ⁠.
- Users should be able to update their profile and view their profile details.
- Admins should be able to create, update, delete, and view user profiles.

2.⁠ ⁠*Design:*
- Design the API endpoints and their methods. Include routes, request types, and responses.
- Explain how you would apply SOLID principles to ensure the API's codebase is maintainable and scalable. Specifically, discuss how you would implement the Single Responsibility Principle (SRP) and Dependency Inversion Principle (DIP) in your API design.
- Identify and describe any design patterns you would use in your API implementation (e.g., Repository Pattern, Service Layer Pattern).
- How would you handle errors and validation in the API?

3.⁠ ⁠*Example:*
- Provide an example of a request and response for one of the API endpoints (e.g., retrieving a user profile).

---

*Answer Outline:*

1.⁠ ⁠*Endpoints and Methods:*
- ⁠ POST /users ⁠ - Create a new user (accessible by admins)
- ⁠ GET /users/{id} ⁠ - Retrieve a user's profile (accessible by the user and admins)
- ⁠ PUT /users/{id} ⁠ - Update a user's profile (accessible by the user and admins)
- ⁠ DELETE /users/{id} ⁠ - Delete a user (accessible by admins)

2.⁠ ⁠*SOLID Principles:*
- *Single Responsibility Principle (SRP):*
    - Each class or module should have one reason to change. For instance, separate classes for handling user requests, user data persistence, and user validation.
- *Dependency Inversion Principle (DIP):*
    - Depend on abstractions (e.g., interfaces) rather than concrete implementations. Use dependency injection to provide dependencies, making it easier to swap out implementations.

3.⁠ ⁠*Design Patterns:*
- *Repository Pattern:* Abstract the data access logic from the business logic. This ensures that changes in data storage (e.g., moving from SQL to NoSQL) do not affect the business logic.
- *Service Layer Pattern:* Encapsulate the business logic in services. This separates business logic from the API layer, making it easier to maintain and test.

4.⁠ ⁠*Error Handling and Validation:*
- Implement centralized error handling using middleware to catch and format errors consistently.
- Validate input data using request validation libraries or frameworks. Ensure meaningful error messages are returned to the client.

5.⁠ ⁠*Example Request and Response:*
- *Request (GET /users/{id}):*
  ⁠ http
  GET /users/123
      ⁠
- *Response:*
  ⁠ json
  {
  "id": 123,
  "name": "John Doe",
  "email": "john.doe@example.com",
  "role": "user"
  }
