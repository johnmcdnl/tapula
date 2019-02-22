package tapula

import (
	"math"
	"sort"
	"time"
)

type Metrics struct {
	Metrics []*Metric `json:"metrics,omitempty"`
}

func (m *Metrics) String() string {
	return toPrettyString(m)
}

func (m *Metrics) Add(others []*Metric) {
	for _, o := range others {
		var foundExisting bool
		for _, m2 := range m.Metrics {
			if m2.Endpoint == o.Endpoint {
				m2.Durations = append(m2.Durations, o.Durations...)
				foundExisting = true
			}
		}
		if !foundExisting {
			m.Metrics = append(m.Metrics, o)
		}
	}
}

type Metric struct {
	Endpoint    string      `json:"endpoint,omitempty"`
	Durations   []*Duration `json:"-"`
	Total       *Duration   `json:"total,omitempty"`
	Count       int         `json:"count,omitempty"`
	Average     *Duration   `json:"average,omitempty"`
	Min         *Duration   `json:"min,omitempty"`
	Max         *Duration   `json:"max,omitempty"`
	P50         *Duration   `json:"p50,omitempty"`
	P75         *Duration   `json:"p75,omitempty"`
	P90         *Duration   `json:"p90,omitempty"`
	P95         *Duration   `json:"p95,omitempty"`
	P99         *Duration   `json:"p99,omitempty"`
	Percentiles *Percentile `json:"-"`
}

func (m *Metric) String() string {
	return toPrettyString(m)
}

func (m *Metric) calculate() {
	m.Count = len(m.Durations)
	m.Max = m.calculateMax()
	m.Min = m.calculateMin()
	m.Total = m.calculateTotal()
	m.Average = m.calculateAverage()

	m.Percentiles = &Percentile{durations: m.Durations}
	m.Percentiles.calculate()
	m.P50 = m.Percentiles.P50
	m.P75 = m.Percentiles.P75
	m.P90 = m.Percentiles.P90
	m.P95 = m.Percentiles.P95
	m.P99 = m.Percentiles.P99

}

func (m *Metric) calculateMax() *Duration {
	max := &Duration{0}
	for _, d := range m.Durations {
		if d.Duration > max.Duration {
			max = d
		}
	}
	return max
}

func (m *Metric) calculateMin() *Duration {
	min := &Duration{math.MaxInt64}
	for _, d := range m.Durations {
		if d.Duration < min.Duration {
			min = d
		}
	}
	return min
}

func (m *Metric) calculateTotal() *Duration {
	sum := &Duration{0}
	for _, d := range m.Durations {
		sum.Duration += d.Duration
	}
	return sum
}

func (m *Metric) calculateAverage() *Duration {
	return &Duration{m.calculateTotal().Duration / time.Duration(len(m.Durations))}
}

func (m *Metric) calculateP(p float64) *Duration {

	values := m.sorted()
	n := float64(len(values))
	i := (p / 100 * n) - 1

	if i != float64(int64(i)) {
		i = math.Ceil(i)
		return values[int(i)]
	}
	return &Duration{(values[int(i)].Duration + values[int(i)+1].Duration) / 2}
}

func (m *Metric) sorted() []*Duration {

	sorted := make([]*Duration, len(m.Durations))
	copy(sorted, m.Durations)

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Duration < sorted[j].Duration
	})

	return sorted
}

type Percentile struct {
	durations []*Duration
	P01       *Duration
	P02       *Duration
	P03       *Duration
	P04       *Duration
	P05       *Duration
	P06       *Duration
	P07       *Duration
	P08       *Duration
	P09       *Duration
	P10       *Duration
	P11       *Duration
	P12       *Duration
	P13       *Duration
	P14       *Duration
	P15       *Duration
	P16       *Duration
	P17       *Duration
	P18       *Duration
	P19       *Duration
	P20       *Duration
	P21       *Duration
	P22       *Duration
	P23       *Duration
	P24       *Duration
	P25       *Duration
	P26       *Duration
	P27       *Duration
	P28       *Duration
	P29       *Duration
	P30       *Duration
	P31       *Duration
	P32       *Duration
	P33       *Duration
	P34       *Duration
	P35       *Duration
	P36       *Duration
	P37       *Duration
	P38       *Duration
	P39       *Duration
	P40       *Duration
	P41       *Duration
	P42       *Duration
	P43       *Duration
	P44       *Duration
	P45       *Duration
	P46       *Duration
	P47       *Duration
	P48       *Duration
	P49       *Duration
	P50       *Duration
	P51       *Duration
	P52       *Duration
	P53       *Duration
	P54       *Duration
	P55       *Duration
	P56       *Duration
	P57       *Duration
	P58       *Duration
	P59       *Duration
	P60       *Duration
	P61       *Duration
	P62       *Duration
	P63       *Duration
	P64       *Duration
	P65       *Duration
	P66       *Duration
	P67       *Duration
	P68       *Duration
	P69       *Duration
	P70       *Duration
	P71       *Duration
	P72       *Duration
	P73       *Duration
	P74       *Duration
	P75       *Duration
	P76       *Duration
	P77       *Duration
	P78       *Duration
	P79       *Duration
	P80       *Duration
	P81       *Duration
	P82       *Duration
	P83       *Duration
	P84       *Duration
	P85       *Duration
	P86       *Duration
	P87       *Duration
	P88       *Duration
	P89       *Duration
	P90       *Duration
	P91       *Duration
	P92       *Duration
	P93       *Duration
	P94       *Duration
	P95       *Duration
	P96       *Duration
	P97       *Duration
	P98       *Duration
	P99       *Duration
}

func (p *Percentile) calculate() {
	p.P01 = p.percentile(1)
	p.P02 = p.percentile(2)
	p.P03 = p.percentile(3)
	p.P04 = p.percentile(4)
	p.P05 = p.percentile(5)
	p.P06 = p.percentile(6)
	p.P07 = p.percentile(7)
	p.P08 = p.percentile(8)
	p.P09 = p.percentile(9)
	p.P10 = p.percentile(10)
	p.P11 = p.percentile(11)
	p.P12 = p.percentile(12)
	p.P13 = p.percentile(13)
	p.P14 = p.percentile(14)
	p.P15 = p.percentile(15)
	p.P16 = p.percentile(16)
	p.P17 = p.percentile(17)
	p.P18 = p.percentile(18)
	p.P19 = p.percentile(19)
	p.P20 = p.percentile(20)
	p.P21 = p.percentile(21)
	p.P22 = p.percentile(22)
	p.P23 = p.percentile(23)
	p.P24 = p.percentile(24)
	p.P25 = p.percentile(25)
	p.P26 = p.percentile(26)
	p.P27 = p.percentile(27)
	p.P28 = p.percentile(28)
	p.P29 = p.percentile(29)
	p.P30 = p.percentile(30)
	p.P31 = p.percentile(31)
	p.P32 = p.percentile(32)
	p.P33 = p.percentile(33)
	p.P34 = p.percentile(34)
	p.P35 = p.percentile(35)
	p.P36 = p.percentile(36)
	p.P37 = p.percentile(37)
	p.P38 = p.percentile(38)
	p.P39 = p.percentile(39)
	p.P40 = p.percentile(40)
	p.P41 = p.percentile(41)
	p.P42 = p.percentile(42)
	p.P43 = p.percentile(43)
	p.P44 = p.percentile(44)
	p.P45 = p.percentile(45)
	p.P46 = p.percentile(46)
	p.P47 = p.percentile(47)
	p.P48 = p.percentile(48)
	p.P49 = p.percentile(49)
	p.P50 = p.percentile(50)
	p.P51 = p.percentile(51)
	p.P52 = p.percentile(52)
	p.P53 = p.percentile(53)
	p.P54 = p.percentile(54)
	p.P55 = p.percentile(55)
	p.P56 = p.percentile(56)
	p.P57 = p.percentile(57)
	p.P58 = p.percentile(58)
	p.P59 = p.percentile(59)
	p.P60 = p.percentile(60)
	p.P61 = p.percentile(61)
	p.P62 = p.percentile(62)
	p.P63 = p.percentile(63)
	p.P64 = p.percentile(64)
	p.P65 = p.percentile(65)
	p.P66 = p.percentile(66)
	p.P67 = p.percentile(67)
	p.P68 = p.percentile(68)
	p.P69 = p.percentile(69)
	p.P70 = p.percentile(70)
	p.P71 = p.percentile(71)
	p.P72 = p.percentile(72)
	p.P73 = p.percentile(73)
	p.P74 = p.percentile(74)
	p.P75 = p.percentile(75)
	p.P76 = p.percentile(76)
	p.P77 = p.percentile(77)
	p.P78 = p.percentile(78)
	p.P79 = p.percentile(79)
	p.P80 = p.percentile(80)
	p.P81 = p.percentile(81)
	p.P82 = p.percentile(82)
	p.P83 = p.percentile(83)
	p.P84 = p.percentile(84)
	p.P85 = p.percentile(85)
	p.P86 = p.percentile(86)
	p.P87 = p.percentile(87)
	p.P88 = p.percentile(88)
	p.P89 = p.percentile(89)
	p.P90 = p.percentile(90)
	p.P91 = p.percentile(91)
	p.P92 = p.percentile(92)
	p.P93 = p.percentile(93)
	p.P94 = p.percentile(94)
	p.P95 = p.percentile(95)
	p.P96 = p.percentile(96)
	p.P97 = p.percentile(97)
	p.P98 = p.percentile(98)
	p.P99 = p.percentile(99)
}

func (p *Percentile) percentile(percentile float64) *Duration {
	sort.Slice(p.durations, func(i, j int) bool {
		return p.durations[i].Duration < p.durations[j].Duration
	})

	n := float64(len(p.durations))
	i := (percentile / 100 * n) - 1

	if i != float64(int64(i)) {
		i = math.Ceil(i)
		return p.durations[int(i)]
	}
	return &Duration{(p.durations[int(i)].Duration + p.durations[int(i)+1].Duration) / 2}
}
