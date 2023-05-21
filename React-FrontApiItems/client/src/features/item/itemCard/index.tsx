import { CardHeader } from "./components/cardHeader";
import { CardContent } from "./components/cardContent";
import { CardText } from "./components/cardText";
import { ItemCard as Card } from "./itemCard";

export const ItemCard = Object.assign(Card, {
    Header: CardHeader,
    Content: CardContent,
    Text: CardText
})