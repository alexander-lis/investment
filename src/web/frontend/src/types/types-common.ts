import {ReactNode} from "react"

// Элемент навигации.
export type NavigationItem = {
    name: string,
    href: string,
    icon: ReactNode,
}

// Начальные данные корневой страницы.
export type LayoutPageData = {
    user: User
}

// Текущий пользователь.
export type User = {
    Name: string
}
