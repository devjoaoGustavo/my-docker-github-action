# My Docker Github action

This action prints a greeting message to the screen

If an argumento is passed through, the greeting will be to the given argument, otherwise, it will be to the world.

## Inputs

### who-to-greet

**Required** string containing the name to be greeted, default to "World".

## Output

The time the greet was sent.

## Example usage

uses: devjoaogustavo/my-docker-github-action@v1
with:
  who-to-greet: 'Fulano'
