# Develop

The simple way to develop is just clone it, make de modification and make a pull request. So is important to note that the modification should be detailed at the pull request, so we can analise it. 

The pull request should explain de follow points:

* Motivation
* What are the gains
* Implementation Details

If you think im make a big modification, you can open a [Issue](https://github.com/gsdenys/pdgen/issues) to talk about, then will be simple to analise The PR, and this issue can be add to the project and planned for an specific release.

Fell free to get a task from any [Project](https://github.com/gsdenys/pdgen/projects?query=is%3Aopen) for you, just comment in the task that you want it.

## Required Development

For the next version we are focused in provide more languages and more export formats. Show bellow in each session how you can contribute for each one of this point.

### Languages

If you want to create a new language follow the steps below:

1. Clone the pdgen repository;
2. Go to the _pkg/services/translate/lang_ and make a copy of _en.go_ renaming to the language that you want to create. e.g. _es.go_;
3. Translate the strings and rename the function to corresponds to your language. e.g. 

```go
func LatinAmericanSpanish(lang language.Tag) language.Tag {
}
```
4. At the file _pkg/services/translate/base.go_ add new language registration for the language at the function _Register_ following the example bellow;

```go
RegLang[language.Spanish.String()] = lang.CanadianFrench(language.Spanish)
RegLang[language.LatinAmericanSpanish.String()] = lang.CanadianFrench(language.LatinAmericanSpanish)
```

5. Make sure that your implementation pass at all unit tests;

?> The threshold for new code coverage is 80%.

6. Make a pull request.


### Formats

If you want to create a new output format (like .docx, .pdf, etc) follow the steps below:

1. Clone the pdgen repository;
2. Go to the _pkg/services/writer_ and create a new _.go_ file to represents the new output format that you want to create;
3. Implement the interface __Printer__ showed below and disponible at _pkg/services/printer.go_. You also can use the _txt.go_ as example;

```go
type Printer interface {
	SetWriter(path string) error
	Init(desc models.Describe)
	Title(title string)
	Subtitle(subtitle string)
	SubSubtitle(subSubtitle string)
	LineBreak()
	Body(desc string)
	Columns(columns []models.Columns)
	Table(t models.Table)
	Done(desc models.Describe)
}
```

4. Register your implementation at __Options__ map variable that are inside the _/Users/gsdenys/Development/gsdenys/pdgen/pkg/options/output.go_ file. 

```go
// Options for output format
var Options = map[string]services.Printer{
	"DEFAULT": &writer.DEFAULT{},
	"MD":      &writer.MD{},
	"TXT":     &writer.TXT{},
	"HTML":    &writer.HTML{},
	"JSON":    &writer.JSON{},
    
    //This is a new one
    "DOCX":     &writer.DOCX{}, 
}
```

5. Make sure that your implementation pass at all unit tests;

?> The threshold for new code coverage is 80%.

6. Make a pull request.
