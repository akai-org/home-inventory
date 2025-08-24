export interface Item {
    id: string;
    name: string;
    type: string;
    description: string;
    image_url: string | null;
    storage_id: string;
    tags: string[];
    created_at: string;
    updated_at: string;
}

export interface Storage {
    id: string;
    name: string;
    parent_id: string | null;
    created_at: string;
    updated_at: string;
}
