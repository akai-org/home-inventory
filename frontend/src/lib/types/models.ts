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
