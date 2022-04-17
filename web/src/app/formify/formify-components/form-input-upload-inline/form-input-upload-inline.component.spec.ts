import { ComponentFixture, TestBed } from '@angular/core/testing';
import { NoopAnimationsModule } from '@angular/platform-browser/animations';
import { FormifyModule } from '../../formify.module';
import { FormInputUploadInlineComponent } from './form-input-upload-inline.component';
import { MockUploadInline } from './upload-inline.mock';


describe('FormInputUploadInlineComponent', () => {
  let component: FormInputUploadInlineComponent;
  let fixture: ComponentFixture<FormInputUploadInlineComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ FormInputUploadInlineComponent ],
      imports: [
        FormifyModule,
        NoopAnimationsModule,
      ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(FormInputUploadInlineComponent);
    component = fixture.componentInstance;
    component.model = new MockUploadInline()
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
