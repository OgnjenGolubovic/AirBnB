export class Accommodation {
    id: string = '';
    name: string = '';
    location: string = '';
    benefits: string = '';
    photos: string = '';
    minGuest: number = 0;
    maxGuest: number = 0;

    public constructor(obj?: any){
        if(obj){
            this.id = obj.id;
            this.name = obj.name;
            this.location = obj.location;
            this.benefits = obj.benefits;
            this.photos = obj.photos;
            this.minGuest = obj.minGuest;
            this.maxGuest = obj.maxGuest;
        }
    }
}