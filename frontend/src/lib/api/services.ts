import type { Item, Storage } from "$lib/types/models";

export async function listItems(params?: URLSearchParams): Promise<Item[]> {
    const url = params ? `/api/v1/items?${params.toString()}` : "/api/v1/items";
    const response = await fetch(url);
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

export async function listStorages(): Promise<Storage[]> {
    const response = await fetch("/api/v1/storages");
    if (!response.ok) {
        throw new Error("Failed to fetch storages");
    }
    return await response.json();
}

export async function createStorage(storage: Storage): Promise<Storage> {
    const response = await fetch("/api/v1/storages", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(storage),
    });
    if (!response.ok) {
        throw new Error("Failed to create storage");
    }
    return await response.json();
}

export async function listStorageItems(storageId: string): Promise<Item[]> {
    const response = await fetch(`/api/v1/storages/${storageId}/items`);
    if (!response.ok) {
        throw new Error("Failed to fetch storage items");
    }
    return await response.json();
}

export async function generateQRCode(entityId: string, entityType: string): Promise<Blob> {
    const response = await fetch("/api/v1/qr/generate", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ entity_id: entityId, entity_type: entityType }),
    });
    if (!response.ok) {
        throw new Error("Failed to generate QR code");
    }
    return await response.blob();
}

export async function generateBarcode(data: string): Promise<Blob> {
    const response = await fetch("/api/v1/barcode/generate", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ data }),
    });
    if (!response.ok) {
        throw new Error("Failed to generate barcode");
    }
    return await response.blob();
}
