package day16

type Packet interface {
	Header() Header
	Value() int
	Inner() []Packet
}

type Header struct {
	Ver int
	Typ int
}

type Literal struct {
	hdr Header
	val int
}

func (l *Literal) Header() Header  { return l.hdr }
func (l *Literal) Value() int      { return l.val }
func (l *Literal) Inner() []Packet { return nil }

type Operator struct {
	hdr    Header
	lenTyp bool // we don't really need this
	inner  []Packet
}

func (o *Operator) Header() Header  { return o.hdr }
func (o *Operator) Inner() []Packet { return o.inner }
func (o *Operator) Value() int {
	var reducer func(p, n int) int
	if len(o.inner) == 1 && o.hdr.Typ >= 0 && o.hdr.Typ <= 3 {
		return o.inner[0].Value()
	}

	switch o.hdr.Typ {
	case 0: // sum
		reducer = func(p, n int) int { return p + n }
	case 1: // product
		reducer = func(p, n int) int { return p * n }
	case 2: // minimum
		reducer = func(p, n int) int {
			if p < n {
				return p
			}
			return n
		}
	case 3: // maximum
		reducer = func(p, n int) int {
			if p > n {
				return p
			}
			return n
		}
	case 5: // greater
		reducer = func(p, n int) int {
			if p > n {
				return 1
			}
			return 0
		}
	case 6: // less than
		reducer = func(p, n int) int {
			if p < n {
				return 1
			}
			return 0
		}
	case 7: // equal
		reducer = func(p, n int) int {
			if p == n {
				return 1
			}
			return 0
		}
	}
	val := reducer(o.inner[0].Value(), o.inner[1].Value())
	for _, n := range o.inner[2:] {
		val = reducer(val, n.Value())
	}
	return val
}

func ParseHeader(b *BitStream) Header {
	var ret Header
	ret.Ver = b.NextInt(3)
	ret.Typ = b.NextInt(3)
	return ret
}

func ParsePacket(b *BitStream) Packet {
	hdr := ParseHeader(b)
	switch hdr.Typ {
	case 4:
		return parseLiteral(b, hdr)
	default:
		return parseOperator(b, hdr)
	}
}

func ParsePackets(b *BitStream) []Packet {
	var ret []Packet
	for b.Len() > 6 {
		ret = append(ret, ParsePacket(b))
	}
	return ret
}

func parseLiteral(b *BitStream, header Header) *Literal {
	ret := &Literal{hdr: header}
	more := true
	for more {
		more = b.NextBool()
		bits := b.NextInt(4)
		ret.val = (ret.val << 4) | bits
	}
	return ret
}

func parseOperator(b *BitStream, header Header) *Operator {
	ret := &Operator{hdr: header, lenTyp: b.NextBool()}
	if ret.lenTyp {
		numPackets := b.NextInt(11)
		for i := 0; i < numPackets; i++ {
			ret.inner = append(ret.inner, ParsePacket(b))
		}
	} else {
		numBits := b.NextInt(15)
		sub := b.Substream(numBits)
		ret.inner = ParsePackets(sub)
	}
	return ret
}

func VersionSum(packets ...Packet) int {
	sum := 0
	for _, p := range packets {
		sum += p.Header().Ver + VersionSum(p.Inner()...)
	}
	return sum
}
