// Инвестиционный портфель.
export type Portfolio = {
    ID: number,
    Name: string,
}

// Группа в рамках инвестиционного портфеля.
export type Group = {
    ID: number,
    Name: string
}

// Уровень группировки в рамках инвестиционного портфеля.
export enum GroupLevel {
    L1 = 1,
    L2 = 2,
    L3 = 3
}

// Инструмент фондового рынка.
export type Instrument = {
    Ticker: string,
}

// Тип инструмента.
export enum InstrumentType {
    Cash = 1,
    Share = 2, 
    GovernmentBond = 3,
    CorporateBond = 4,
}