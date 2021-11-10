### HTML Link Parser:

We are going to parse the HTML and get the all links in that page. We need to return the `href` and `Text` of the Link as a golang struct.
We are going to ignore if there is any nested links in the `Text`. We will strip the extra spaces and newlines in Text.

We will use the https://pkg.go.dev/golang.org/x/net/html package to parse the HTML pages.

### Example:

```
<html>
<body>
  <h1>Hello World!</h1>
  <a href="/test-page">A link to test page</a>
</body>
</html>
```

For this we need to return something like below

```
{
  Href: "/test-page",
  Text: "A link to test page",
}
```

Note: This is my practise version of one of the execerises from *gophercises*

### Example Outputs:

- For example1.html
  ```
  > go run .\main.go
  [{/test-page A link to test page}]
  >
  ```

- For example2.html
  ```
  > go run .\main.go
  [{https://www.nasa.gov/ NASA Website } {https://www.discovery.com/ Discovery Website}]
  >
  ```

- For example3.html
  ```
  > go run .\main.go
  [{# Login} {/lost Lost? Need help?} {https://twitter.com/marcusolsson @marcusolsson}]
  >
  ```

- For example4.html
  ```
  > go run .\main.go
  [{/dog-cat dog cat}]
  >
  ```


