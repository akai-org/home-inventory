import type { Item } from "$lib/types/models";

export async function listItems(): Promise<Item[]> {
    const response = await fetch("/api/v1/items");
    if (!response.ok) {
        throw new Error("Failed to fetch items");
    }
    return await response.json();
}

export async function createItem(item: Item): Promise<Item> {
    const response = await fetch("/api/v1/items", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(item),
    });
    if (!response.ok) {
        throw new Error("Failed to create item");
    }
    return await response.json();
}
