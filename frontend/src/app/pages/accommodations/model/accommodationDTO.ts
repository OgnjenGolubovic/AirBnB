export interface AccommodationDTO {
    id: string;
    name: string;
    location: string;
    benefits: string;
    photos: string;
    hostId: string;
    minGuest: number;
    maxGuest: number;
    automaticApproval: boolean;
    price: number;
    dates: string[];
    isPerGuest: boolean;
    hasWeekend: boolean;
    hasSummer: boolean;
}