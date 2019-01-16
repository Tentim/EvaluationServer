import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import {DropdownModule} from 'primeng/primeng';
import { AppComponent } from './app.component';
import {LoginComponent, MyModel} from './login/login.component';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    MyModel,
  ],
  imports: [
    BrowserModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
