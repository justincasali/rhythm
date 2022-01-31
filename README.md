# rhythm
[Euclidean Rhythm](https://en.wikipedia.org/wiki/Euclidean_rhythm) generator written in Go with nested [circular lists](https://pkg.go.dev/container/ring) 🤹

## Usage
```
rhythm [hits] [steps] [shift]
```

- `hits` -- number of hits in output sequence
- `steps` -- number of steps in output sequence
- `shift` -- degree of shift in output sequence

## Example
```
rhythm 17 24 2
```

```
[x x x . x x . x x x . x x . x x x . x x . x x .]
```
