# go-smmry

API client for [SMMRY](http://smmry.com/)

## Requirements

Sign up and get an API key from the [SMMRY](http://smmry.com/) website, then set your env variable accordingly.

```bash
export SMMRY_API_KEY="YOUR_SMMRY_API_KEY"
```

## Examples

### Summary by Website

```
go run cmd/bywebsite/bywebsite.go [website_url] [summary_length]
```

e.g.

```
go run cmd/bywebsite/bywebsite.go https://en.wikipedia.org/wiki/Richard_Stallman 5
```

```
2016/10/15 15:11:50 Character count: 1309
2016/10/15 15:11:50 Title: Richard Stallman
2016/10/15 15:11:50 Summary:  Richard Matthew Stallman, often known by his initials, rms, is an American software freedom activist and programmer. Stallman announced the plan for the GNU operating system in September 1983 on several ARPANET mailing lists and USENET. Stallman started the project on his own and describes: "As an operating system developer, I had the right skills for this job. So even though I could not take success for granted, I realized that I was elected to do the job. I chose to make the system compatible with Unix so that it would be portable, and so that Unix users could easily switch to it." Stallman's staunch advocacy for free software inspired the creation of the Virtual Richard M. Stallman, software that analyzes the packages currently installed on a Debian GNU/Linux system, and reports those that are from the non-free tree. Protesting against proprietary software in April 2006, Stallman held a "Don't buy from ATI, enemy of your freedom" placard at a speech by an ATI representative in the building where Stallman worked, resulting in the police being called. After initially accepting the concept, Stallman rejects a common alternative term, open source software, because it does not call to mind what Stallman sees as the value of the software: freedom.
2016/10/15 15:11:56 Quota info: Waited 0 extra seconds due to API Free mode, 74 requests left to make for today.
```

### Summary by Text

```
go run cmd/bytext/bytext.go [text] [summary_length]
```

e.g.
```
go run cmd/bytext/bytext.go "Donald John Trump (born June 14, 1946) is an American businessman, television producer, author, politician, and the Republican Party nominee for President of the United States in the 2016 election. He is the chairman and president of The Trump Organization, which is the principal holding company for his real estate ventures and other business interests. During his career, Trump has built office towers,hotels, casinos, golf courses, an urban development project in Manhattan, and other branded facilities worldwide. Born and raised in the borough of Queens in New York City, Trump received a bachelor's degree in economics from the WhartonSchool of the University of Pennsylvania in 1968. While attending college, he worked in his father Fred Trump's real estate and construction firm. He was given control of the business in 1971 and later renamed it The Trump Organization. Trump has made cameo appearances in films and television series, and he appeared at the Miss USA pageants, which he owned from 1996 to 2015. He sought the Reform Party presidential nomination in 2000, but withdrew before voting began. He hosted and co-produced The Apprentice, a reality television series on NBC, from 2004 to 2015. Trump and his businesses, as well as his personal life and political views, have for decades received considerable media exposure. He is listed byForbes among the world's wealthiest 500 billionaires. In June 2015, Trump announced his candidacy for president as a Republican and quickly emerged as the front-runner for his party's nomination. In May 2016, his remaining Republican rivals suspended their campaigns, and in July he was formally nominated for president at the 2016 Republican National Convention. Trump's campaign has received extensive media coverage and international attention. Many of his statements ininterviews, on Twitter, and at campaign rallies have been controversial or false. Several rallies during the primaries were accompanied by protests or riots. Trump's positions include renegotiation of U.S.–China trade deals, opposition to particular trade agreements such as NAFTA and the Trans-Pacific Partnership, stronger enforcement of immigration laws together with building a wall along the U.S.–Mexico border, reform of veterans' care, repeal and replacement of the Affordable Care Act, and tax cuts. During the primary, Trump called for a temporary ban on all Muslim immigration to the United States, later stating that the ban would focus instead on countries with a proven history of terrorism, until the level of vetting can be raised to screen out potential terrorists. During the 2016 presidential campaign, numerous allegations of sexual misconduct by Trump surfaced, in the aftermath of the release of a 2005 video in which Trump allegedlydescribed sexually assaulting women. In the most serious cases Trump has been accused of rape, sexual assault and child sexual abuse dated from the 1980s to 2010s. He denies the allegations, describing them as a smear campaign." 5
```

```
2016/10/15 15:14:13 Character count: 766
2016/10/15 15:14:13 Title:
2016/10/15 15:14:13 Summary:  Donald John Trump is an American businessman, television producer, author, politician, and the Republican Party nominee for President of the United States in the 2016 election. During his career, Trump has built office towers,hotels, casinos, golf courses, an urban development project in Manhattan, and other branded facilities worldwide. In June 2015, Trump announced his candidacy for president as a Republican and quickly emerged as the front-runner for his party's nomination. Trump's campaign has received extensive media coverage and international attention. During the 2016 presidential campaign, numerous allegations of sexual misconduct by Trump surfaced, in the aftermath of the release of a 2005 video in which Trump allegedly described sexually assaulting women.
2016/10/15 15:14:13 Quota info: Waited 0 extra seconds due to API Free mode, 73 requests left to make for today.
```

## License

MIT
