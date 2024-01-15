// Начальные данные страницы портфелей.
export type PortfoliosPageData = {
    portfolios: PortfolioLink[]
}

// Начальные данные страницы портфеля.
export type PortfolioPageData = {
    portfolio: Portfolio,
    groups: Group[]
    sectors: Sector[]
    assets: Asset[]
}

// Ссылки на инвестиционный портфель.
export type PortfolioLink = {
    id: number,
    name: string,
}

// Полные данные инвестиционного портфеля.
export type Portfolio = {
    id: number,
    name: string,
    startDate: Date,
    endDate: Date
}

// Группа рисков актива.
export type Group = {
    id: number,
    name: string,
}

// Сектор актива.
export type Sector = {
    id: number,
    groupId: number,
    name: string,
}

// Инструмент фондового рынка.
export type Asset = {
    id: number,
    sectorId: number,
    name: string,
    ticker: string,
    price: number,
}

// Тип инструмента.
export enum AssetType {
    Cash = 1,
    Share = 2, 
    Bond = 3,
}