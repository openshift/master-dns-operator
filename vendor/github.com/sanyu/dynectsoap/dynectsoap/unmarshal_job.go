package dynectsoap

import (
	"encoding/xml"
	"errors"
	"io"
	"regexp"
)

func (c *GetJobResponseType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var errorResponse = regexp.MustCompile(`.*ErrorResponseType`)
	var getAllRecordsResponse = regexp.MustCompile(`.*GetAllRecordsResponseType`)

	for {
		token, err := d.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		var decodeErr error
		switch token.(type) {
		case xml.StartElement:
			nextStart := token.(xml.StartElement)
			local := nextStart.Name.Local
			switch local {
			case "msgs":
				decodeErr = d.DecodeElement(&c.Msgs, &nextStart)
			case "status":
				decodeErr = d.DecodeElement(&c.Status, &nextStart)
			case "job_id":
				decodeErr = d.DecodeElement(&c.Job_id, &nextStart)
			case "data":
				for _, attr := range nextStart.Attr {
					if attr.Name.Local == "type" {
						responseType := attr.Value
						switch {
						case errorResponse.MatchString(responseType):
							c.Data = new(ErrorResponseType)
							decodeErr = d.DecodeElement(&c.Data, &nextStart)
						case getAllRecordsResponse.MatchString(responseType):
							c.Data = new(GetAllRecordsResponseType)
							decodeErr = d.DecodeElement(&c.Data, &nextStart)
						default:
							decodeErr = errors.New("Unknown node")
						}
					}
				}

			}
		}
		if decodeErr != nil {
			return decodeErr
		}
	}

	return nil
}
