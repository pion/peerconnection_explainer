// Code generated by "json-ice --type=MediaSectionDetails"; DO NOT EDIT.

package result

func MarshalMediaSectionDetailsAsJSON(s *MediaSectionDetails) ([]byte, error) {
	buff := make([]byte, 1, 132)
	buff[0] = '{'
	buff = append(buff, "\"ID\":"...)
	if marshaled, err := s.ID.MarshalJSON(); err != nil {
		return nil, err
	} else {
		buff = append(buff, marshaled...)
	}

	buff = append(buff, ',')
	buff = append(buff, "\"type\":"...)
	if marshaled, err := s.Type.MarshalJSON(); err != nil {
		return nil, err
	} else {
		buff = append(buff, marshaled...)
	}

	buff = append(buff, ',')
	buff = append(buff, "\"direction\":"...)
	if marshaled, err := s.Direction.MarshalJSON(); err != nil {
		return nil, err
	} else {
		buff = append(buff, marshaled...)
	}

	buff = append(buff, ',')
	if s.MediaFormats == nil {
		buff = append(buff, "\"mediaFormats\":null,"...)
	} else {
		buff = append(buff, "\"mediaFormats\":"...)
		buff = append(buff, '[')
		for _, v := range s.MediaFormats {
			if marshaled, err := v.MarshalJSON(); err != nil {
				return nil, err
			} else {
				buff = append(buff, marshaled...)
			}

			buff = append(buff, ',')
		}
		if buff[len(buff)-1] == ',' {
			buff[len(buff)-1] = ']'
		} else {
			buff = append(buff, ']')
		}

		buff = append(buff, ',')
	}
	if buff[len(buff)-1] == ',' {
		buff[len(buff)-1] = '}'
	} else {
		buff = append(buff, '}')
	}
	return buff, nil
}