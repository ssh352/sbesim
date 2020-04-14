// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"math"
)

type MaturityMonthYear struct {
	Year  uint16
	Month uint8
	Day   uint8
	Week  uint8
}

func (m *MaturityMonthYear) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint16(_w, m.Year); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, m.Month); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, m.Day); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, m.Week); err != nil {
		return err
	}
	return nil
}

func (m *MaturityMonthYear) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if !m.YearInActingVersion(actingVersion) {
		m.Year = m.YearNullValue()
	} else {
		if err := _m.ReadUint16(_r, &m.Year); err != nil {
			return err
		}
	}
	if !m.MonthInActingVersion(actingVersion) {
		m.Month = m.MonthNullValue()
	} else {
		if err := _m.ReadUint8(_r, &m.Month); err != nil {
			return err
		}
	}
	if !m.DayInActingVersion(actingVersion) {
		m.Day = m.DayNullValue()
	} else {
		if err := _m.ReadUint8(_r, &m.Day); err != nil {
			return err
		}
	}
	if !m.WeekInActingVersion(actingVersion) {
		m.Week = m.WeekNullValue()
	} else {
		if err := _m.ReadUint8(_r, &m.Week); err != nil {
			return err
		}
	}
	return nil
}

func (m *MaturityMonthYear) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.YearInActingVersion(actingVersion) {
		if m.Year != m.YearNullValue() && (m.Year < m.YearMinValue() || m.Year > m.YearMaxValue()) {
			return fmt.Errorf("Range check failed on m.Year (%v < %v > %v)", m.YearMinValue(), m.Year, m.YearMaxValue())
		}
	}
	if m.MonthInActingVersion(actingVersion) {
		if m.Month != m.MonthNullValue() && (m.Month < m.MonthMinValue() || m.Month > m.MonthMaxValue()) {
			return fmt.Errorf("Range check failed on m.Month (%v < %v > %v)", m.MonthMinValue(), m.Month, m.MonthMaxValue())
		}
	}
	if m.DayInActingVersion(actingVersion) {
		if m.Day != m.DayNullValue() && (m.Day < m.DayMinValue() || m.Day > m.DayMaxValue()) {
			return fmt.Errorf("Range check failed on m.Day (%v < %v > %v)", m.DayMinValue(), m.Day, m.DayMaxValue())
		}
	}
	if m.WeekInActingVersion(actingVersion) {
		if m.Week != m.WeekNullValue() && (m.Week < m.WeekMinValue() || m.Week > m.WeekMaxValue()) {
			return fmt.Errorf("Range check failed on m.Week (%v < %v > %v)", m.WeekMinValue(), m.Week, m.WeekMaxValue())
		}
	}
	return nil
}

func MaturityMonthYearInit(m *MaturityMonthYear) {
	m.Year = 65535
	m.Month = 255
	m.Day = 255
	m.Week = 255
	return
}

func (*MaturityMonthYear) EncodedLength() int64 {
	return 5
}

func (*MaturityMonthYear) YearMinValue() uint16 {
	return 0
}

func (*MaturityMonthYear) YearMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*MaturityMonthYear) YearNullValue() uint16 {
	return 65535
}

func (*MaturityMonthYear) YearSinceVersion() uint16 {
	return 0
}

func (m *MaturityMonthYear) YearInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.YearSinceVersion()
}

func (*MaturityMonthYear) YearDeprecated() uint16 {
	return 0
}

func (*MaturityMonthYear) MonthMinValue() uint8 {
	return 0
}

func (*MaturityMonthYear) MonthMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*MaturityMonthYear) MonthNullValue() uint8 {
	return 255
}

func (*MaturityMonthYear) MonthSinceVersion() uint16 {
	return 0
}

func (m *MaturityMonthYear) MonthInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MonthSinceVersion()
}

func (*MaturityMonthYear) MonthDeprecated() uint16 {
	return 0
}

func (*MaturityMonthYear) DayMinValue() uint8 {
	return 0
}

func (*MaturityMonthYear) DayMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*MaturityMonthYear) DayNullValue() uint8 {
	return 255
}

func (*MaturityMonthYear) DaySinceVersion() uint16 {
	return 0
}

func (m *MaturityMonthYear) DayInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.DaySinceVersion()
}

func (*MaturityMonthYear) DayDeprecated() uint16 {
	return 0
}

func (*MaturityMonthYear) WeekMinValue() uint8 {
	return 0
}

func (*MaturityMonthYear) WeekMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*MaturityMonthYear) WeekNullValue() uint8 {
	return 255
}

func (*MaturityMonthYear) WeekSinceVersion() uint16 {
	return 0
}

func (m *MaturityMonthYear) WeekInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.WeekSinceVersion()
}

func (*MaturityMonthYear) WeekDeprecated() uint16 {
	return 0
}
