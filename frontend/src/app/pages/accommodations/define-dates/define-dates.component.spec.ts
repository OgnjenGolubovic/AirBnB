import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DefineDatesComponent } from './define-dates.component';

describe('DefineDatesComponent', () => {
  let component: DefineDatesComponent;
  let fixture: ComponentFixture<DefineDatesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DefineDatesComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(DefineDatesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
