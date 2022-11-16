# Input Validation in rk_echo

## Tutorials

- Input validation can be easily done as follows:

```golang
type CreateUserInput struct {
	Name string `json:"name" validate:"required"`
}

func CreateUser(ctx echo.Context) error {
	// Get the input body from request
    input := new(CreateUserInput)   // Init input object
    err := echoutils.EchoGetInputAndValidate(ctx, input)    // Validate and parse input in request to that obj
	
	// If there is any error during validation, return it
    if err != nil {
        return err.ReturnErrorMessage(ctx)
    }
    newUser, err := CreateUserLogic(input.Name)
    if err != nil {
        return err.ReturnErrorMessage(ctx)
    }
    return ctx.JSON(http.StatusOK, echoutils.ResponseMessage{
        Message: "Create user successfully!!",
        Data:    newUser,
    })
}
```

- For validation syntax, please check this link here to learn more:
    - https://github.com/go-playground/validator

