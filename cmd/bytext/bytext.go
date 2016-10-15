package main

/*
Example:
go run cmd/bytext/bytext.go "Donald John Trump (born June 14, 1946) is an American businessman, television producer, author, politician, and the Republican Party nominee for President of the United States in the 2016 election. He is the chairman and president of The Trump Organization, which is the principal holding company for his real estate ventures and other business interests. During his career, Trump has built office towers,hotels, casinos, golf courses, an urban development project in Manhattan, and other branded facilities worldwide. Born and raised in the borough of Queens in New York City, Trump received a bachelor's degree in economics from the WhartonSchool of the University of Pennsylvania in 1968. While attending college, he worked in his father Fred Trump's real estate and construction firm. He was given control of the business in 1971 and later renamed it The Trump Organization. Trump has made cameo appearances in films and television series, and he appeared at the Miss USA pageants, which he owned from 1996 to 2015. He sought the Reform Party presidential nomination in 2000, but withdrew before voting began. He hosted and co-produced The Apprentice, a reality television series on NBC, from 2004 to 2015. Trump and his businesses, as well as his personal life and political views, have for decades received considerable media exposure. He is listed by Forbes among the world's wealthiest 500 billionaires. In June 2015, Trump announced his candidacy for president as a Republican and quickly emerged as the front-runner for his party's nomination. In May 2016, his remaining Republican rivals suspended their campaigns, and in July he was formally nominated for president at the 2016 Republican National Convention. Trump's campaign has received extensive media coverage and international attention. Many of his statements ininterviews, on Twitter, and at campaign rallies have been controversial or false. Several rallies during the primaries were accompanied by protests or riots. Trump's positions include renegotiation of U.S.–China trade deals, opposition to particular trade agreements such as NAFTA and the Trans-Pacific Partnership, stronger enforcement of immigration laws together with building a wall along the U.S.–Mexico border, reform of veterans' care, repeal and replacement of the Affordable Care Act, and tax cuts. During the primary, Trump called for a temporary ban on all Muslim immigration to the United States, later stating that the ban would focus instead on countries with a proven history of terrorism, until the level of vetting can be raised to screen out potential terrorists. During the 2016 presidential campaign, numerous allegations of sexual misconduct by Trump surfaced, in the aftermath of the release of a 2005 video in which Trump allegedly described sexually assaulting women. In the most serious cases Trump has been accused of rape, sexual assault and child sexual abuse dated from the 1980s to 2010s. He denies the allegations, describing them as a smear campaign." 3
*/

import (
	"log"
	"os"

	"github.com/microamp/go-smmry/smmry"
)

func main() {
	args := os.Args[1:]
	text, length := args[0], args[1]

	client, err := smmry.NewSmmryClient()
	if err != nil {
		log.Panicf("Error instantiating SMMRY client: %+v", err)
	}

	summary, err := client.SummaryByText(text, length)
	if err != nil {
		log.Panicf("Error getting summary by website: %+v", err)
	}

	log.Printf("Character count: %s", summary.SmAPICharacterCount)
	log.Printf("Title: %s", summary.SmAPITitle)
	log.Printf("Summary: %s", summary.SmAPIContent)
	log.Printf("Quota info: %s", summary.SmAPILimitation)
}
