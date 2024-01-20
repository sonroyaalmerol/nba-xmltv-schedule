package xmltv

import "encoding/xml"

// Program represents a single program entry
type Program struct {
	ID         string     `xml:"id,attr,omitempty"`
	Start      string     `xml:"start,attr"`
	Stop       string     `xml:"stop,attr"`
	Channel    string     `xml:"channel,attr"`
	Title      Title      `xml:"title"`
	SubTitle   SubTitle   `xml:"sub-title,omitempty"`
	Desc       string     `xml:"desc,omitempty"`
	Date       string     `xml:"date,omitempty"`
	Categories []Category `xml:"category,omitempty"`
	EpisodeNum string     `xml:"episode-num,omitempty"`
	Audio      Audio      `xml:"audio,omitempty"`
	Previously *TimeSpan  `xml:"previously-shown,omitempty"`
	Subtitles  Subtitles  `xml:"subtitles,omitempty"`
	Rating     *Rating    `xml:"rating,omitempty"`
}

type Title struct {
	Lang string `xml:"lang,attr"`
	Text string `xml:",chardata"`
}

type Category struct {
	Lang string `xml:"lang,attr"`
	Text string `xml:",chardata"`
}

type SubTitle struct {
	Lang string `xml:"lang,attr"`
	Text string `xml:",chardata"`
}

// Audio represents audio information
type Audio struct {
	Stereo string `xml:"stereo"`
}

// TimeSpan represents a time span
type TimeSpan struct {
	Start string `xml:"start,attr"`
}

// Subtitles represents subtitles information
type Subtitles struct {
	Type string `xml:"type,attr"`
}

// Rating represents rating information
type Rating struct {
	System string `xml:"system,attr"`
	Value  string `xml:"value"`
}

// TV represents the main TV schedule structure
type TV struct {
	XMLName           xml.Name  `xml:"tv"`
	Channels          []Channel `xml:"channel"`
	Programs          []Program `xml:"programme"`
	SourceInfoUrl     string    `xml:"source-info-url,attr"`
	SourceInfoName    string    `xml:"source-info-name,attr"`
	GeneratorInfoName string    `xml:"generator-info-name,attr"`
	GeneratorInfoUrl  string    `xml:"generator-info-url,attr"`
}

// Channel represents a TV channel entry
type Channel struct {
	ID          string `xml:"id,attr"`
	DisplayName string `xml:"display-name"`
	Icon        Icon   `xml:"icon"`
}

// Icon represents icon information
type Icon struct {
	Src string `xml:"src,attr"`
}
