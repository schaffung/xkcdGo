# xkcdGo
A project aimed towards querying and getting XKCD comic transcripts.

Now one can query the XKCD website using the url -> `https://xkcd/<num>/info.0.json`, where num is the comic number
The resultant is a json of the following form

```javascript
  ` {
      "month"      : "<month_value>",
      "num"        : "<num_value>",
      "link"       : "<link_value>",
      "year"       : "<year_value>",
      "news"       : "<news_value>",
      "safe_title" : "<safe_title_value>",
      "alt"        : "<alternative_text>",
      "transcript" : "<transcript_value>",
      "img"        : "<The image link>",
      "day"        : "<day_value>"
    }`
```

One can import the xkcdGo package and then invoke the SearchComic function to get a resultant struct which will contain all the above data.
