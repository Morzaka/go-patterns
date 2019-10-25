// Pattern Specification
//
// У наступному прикладі ми витягуємо рахунки-фактури та надсилаємо їх до інкасаційного агентства, якщо
//  1. вони прострочені,
//  2. повідомлення надіслані,
//  3. вони ще не в колекторському агентстві.
// Цей приклад має на меті показати кінцевий результат того, як логіка є «пов'язана» між собою.
//
//У цьому прикладі передбачається наявність раніше визначеної
// структури OverdueSpecification,
//  яка задовольняється, коли термін оплати рахунку-фактури складає 30 днів або більше;
// структури NoticeSentSpecification, яка задовольняється, коли клієнту відправляються три повідомлення;
// структури InCollectionSpecification, яка задовольняється,
//  коли рахунок-фактура має вже відправлена в агентство по збору платежів.

package Specification

// Дані для аналізу
type Invoice struct {
	Day    int
	Notice int
	IsSent   bool
}

// Інтерфейс специфікації рахунків-фактур
type Specification interface {
	IsSatisfiedBy(Invoice) bool
	And(Specification) Specification
	Or(Specification) Specification
	Not() Specification
	Relate(Specification)
}

// Базовий рахунок-фактура
type BaseSpecification struct {
	Specification
}
// Перевірка специфікації
func (selfy *BaseSpecification) IsSatisfiedBy(elm Invoice) bool {
	return false
}
func (selfy *BaseSpecification) And(spec Specification) Specification {
	a := &AndSpecification{
		selfy.Specification, spec,
	}
	a.Relate(a)
	return a
}
func (selfy *BaseSpecification) Or(spec Specification) Specification {
	a := &OrSpecification{
		selfy.Specification, spec,
	}
	a.Relate(a)
	return selfy
}
func (selfy *BaseSpecification) Not() Specification {
	a := &NotSpecification{
		selfy.Specification,
	}
	a.Relate(a)
	return a
}
func (selfy *BaseSpecification) Relate(spec Specification) {
	selfy.Specification = spec
}

///////////////////////

type AndSpecification struct {
	Specification
	compare Specification
}
func (sa *AndSpecification) IsSatisfiedBy(elm Invoice) bool {
	return sa.Specification.IsSatisfiedBy(elm) && sa.compare.IsSatisfiedBy(elm)
}

type OrSpecification struct {
	Specification
	compare Specification
}
func (so *OrSpecification) IsSatisfiedBy(elm Invoice) bool {
	return so.Specification.IsSatisfiedBy(elm) || so.compare.IsSatisfiedBy(elm)
}

type NotSpecification struct {
	Specification
}
func (sn *NotSpecification) IsSatisfiedBy(elm Invoice) bool {
	return !sn.Specification.IsSatisfiedBy(elm)
}

//////////////////////////
// Термін сплати рахунка-фактури - 30 днів або старше
type OverDueSpecification struct {
	Specification
}
func (od *OverDueSpecification) IsSatisfiedBy(elm Invoice) bool {
	return elm.Day >= 30
}
func NewOverDueSpecification() Specification {
	a := &OverDueSpecification{&BaseSpecification{}}
	a.Relate(a)
	return a
}

// Три сповіщення надіслані боржнику
type NoticeSendSpecification struct {
	Specification
}
func( nos *NoticeSendSpecification) IsSatisfiedBy(elm Invoice) bool {
	return elm.Notice >= 3
}
func NewNoticeSpecification() Specification {
	a := &NoticeSendSpecification{&BaseSpecification{}}
	a.Relate(a)
	return a
}

// Рахунок-фактура вже надісланий в інкасаційне агентство.
type InCollectionSpecification struct {
	Specification
}
func (self *InCollectionSpecification) IsSatisfiedBy(elm Invoice) bool {
	return !elm.IsSent
}
func NewInCollectionSpecification() Specification {
	a := &InCollectionSpecification{&BaseSpecification{}}
	a.Relate(a)
	return a
}