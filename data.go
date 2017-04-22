package harvester

type Seeds map[MetadataKey]string

func (s Seeds) ToData() Data {
	result := make(Data, len(s))
	for k, v := range s {
		result.Add(k, v, "seed")
	}

	return result
}

type Data map[MetadataKey][]*Metadata

func (d Data) Count() int {
	count := 0
	for _, metas := range d {
		for range metas {
			count++
		}
	}

	return count
}

func (d Data) ForEach(k MetadataKey, f func(m *Metadata) error) error {
	for _, meta := range d[k] {
		if err := f(meta); err != nil {
			return err
		}
	}

	return nil
}

func (d Data) Add(k MetadataKey, v string, pitchfork string) {
	if v != "" {
		if _, ok := d[k]; ok {
			for _, metadata := range d[k] {
				if metadata.Value == v {
					return
				}
			}
		}

		d[k] = append(d[k], &Metadata{
			Value:     v,
			Pitchfork: pitchfork,
		})
	}
}
