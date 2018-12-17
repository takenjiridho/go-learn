import { Component, Input, ElementRef, ViewChild, HostListener  } from '@angular/core';
import {NgbTooltip, NgbTooltipConfig} from '@ng-bootstrap/ng-bootstrap';

export enum KEY_CODE {
  ESC = 27
}

@Component({
  selector: 'app-tooltipper',
  templateUrl: './tooltipper.component.html',
  styleUrls: ['./tooltipper.component.scss'],
  providers: [NgbTooltip, NgbTooltipConfig] // add NgbTooltipConfig to the component providers
})
export class TooltipperComponent {

  constructor(config: NgbTooltip, ngbTooltipConfig: NgbTooltipConfig, private _eref: ElementRef) {
    ngbTooltipConfig.placement = 'auto';
  }

  @Input() info1: string;
  @Input() info2: string;

  @HostListener('window:keyup', ['$event'])

  keyEvent(event: KeyboardEvent) {
    if (event.keyCode === KEY_CODE.ESC) {
      this.tooltip1.close();
      this.tooltip2.close();
    }
  }

  info = {};

  @ViewChild('t') public tooltip1: NgbTooltip;
  @ViewChild('y') public tooltip2: NgbTooltip;

  onClickedOutside(e: Event) {
    this.tooltip2.close();
    this.tooltip1.close();
  }

  public showInfo(info: any): void {
    let isOpen = false;

    if (info.name === 'A') {
      isOpen = this.tooltip1.isOpen();
      this.tooltip2.close();

      if (info !== this.info || !isOpen) {
        info['info'] = this.info1;
        this.info = info['info'];
        this.tooltip1.open(info);
      }
    }

    if (info.name === 'B') {
      isOpen = this.tooltip2.isOpen();
      this.tooltip1.close();

      if (info !== this.info || !isOpen) {
        info['info'] = this.info2;
        this.info = info['info'];
        this.tooltip2.open(info);
      }
    }
  }
}
