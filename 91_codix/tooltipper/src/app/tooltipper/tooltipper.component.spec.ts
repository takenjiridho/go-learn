import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { TooltipperComponent } from './tooltipper.component';

describe('TooltipperComponent', () => {
  let component: TooltipperComponent;
  let fixture: ComponentFixture<TooltipperComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ TooltipperComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(TooltipperComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
