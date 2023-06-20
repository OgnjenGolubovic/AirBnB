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

export interface DatesRange {
    startDate: string;
    endDate: string;
}

export interface DateRequest {
    id: string;
    startDate: string;
    endDate: string;
}