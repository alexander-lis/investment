import {LayoutPageData, User} from "../types/types-common";
import {
    PortfolioLink,
    PortfolioPageData,
    PortfoliosPageData,
} from "../types/types-portfolio";
import {createPromiseClient} from "@connectrpc/connect";
import {createGrpcWebTransport} from "@connectrpc/connect-web";
import {PortfolioService} from "../connect-rpc/services/stock/proto/v1/portfolios_connect.ts";

// Main data.
export async function layoutPageLoader(): Promise<LayoutPageData> {
    const user: User = {
        Name: "alexander",
    };
    return {user};
}

// Portfolios list page.
export async function portfoliosPageLoader(): Promise<PortfoliosPageData> {
    const transport = createGrpcWebTransport({
        baseUrl: "https://127.0.0.1:59691",
    });

    const portfolioClient = createPromiseClient(PortfolioService, transport)
    const portfoliosResponse = await portfolioClient.getPortfolios({})

    const portfolios: PortfolioLink[] = []
    portfoliosResponse.Portfolios.map(p => {
        portfolios.push({
            id: p.Id,
            name: p.Name
        })
    })

    return {portfolios};
}

// Portfolio planner page.
export async function portfolioPageLoader({
                                              params,
                                          }: {
    params: {
        portfolioId?: number;
    };
}): Promise<PortfolioPageData> {
    if (params.portfolioId == 1) {
        return {
            portfolio: {
                id: 1,
                name: "На пенсию",
                startDate: new Date(2020, 9, 10),
                endDate: new Date(2054, 11, 14),
            },
            groups: [
                {
                    id: 1,
                    name: "Консервативные",
                },
                {
                    id: 2,
                    name: "Умеренные",
                },
                {
                    id: 3,
                    name: "Агрессивные",
                },
            ],
            sectors: [
                {
                    id: 1,
                    name: "ОФЗ",
                    groupId: 1,
                },
                {
                    id: 2,
                    name: "Корп",
                    groupId: 1,
                },
                {
                    id: 3,
                    name: "Нефть",
                    groupId: 2,
                },
                {
                    id: 4,
                    name: "IT",
                    groupId: 2,
                },
                {
                    id: 5,
                    name: "Финансы",
                    groupId: 2,
                },
            ],
            assets: [{
                id: 1,
                sectorId: 1,
                name: "ОФЗ 32300",
                ticker: "OFAH",
                price: 20000
            }, {
                id: 2,
                sectorId: 1,
                name: "ОФЗ 24500",
                ticker: "OFSS",
                price: 23000
            }, {
                id: 3,
                sectorId: 2,
                name: "Самолет-1 выпуск 23233",
                ticker: "SMLL1",
                price: 15000
            }],
        };
    }

    if (params.portfolioId == 2) {
        return {
            portfolio: {
                id: 2,
                name: "На квартиру",
                startDate: new Date(2023, 7, 25),
                endDate: new Date(2025, 5, 31),
            },
            groups: [],
            sectors: [],
            assets: [],
        };
    }

    return {
        portfolio: {
            id: 0,
            name: "Ничего не найдено",
            startDate: new Date(),
            endDate: new Date(),
        },
        groups: [],
        sectors: [],
        assets: [],
    };
}

// sum.js
export function sum(a: number, b: number) {
    return a + b;
}
