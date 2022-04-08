import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { AddImageComponent } from 'src/app/forms/add-image/add-image.component';
import { ImageItem } from 'src/types';
import { ApiService } from '../../services/api.service';
import { ListColumns } from '../columns';
import { ImagesTableDataSource } from './images-table-datasource';

@Component({
  selector: 'app-images-list',
  templateUrl: './images-list.component.html',
  styleUrls: ['./images-list.component.scss']
})
export class ImagesListComponent implements OnInit {
  dataSource: ImagesTableDataSource
  columns = ListColumns.imagesColumns

  constructor(private apiService: ApiService, public dialog: MatDialog) {
    this.dataSource = new ImagesTableDataSource(apiService)
  }

  editImage = (image: ImageItem) => {
    this.openEditImageDialog(image)
  }

  addImage = () => {
    this.openAddImageDialog()
  }

  openEditImageDialog(data: ImageItem) {
    // const dialogRef = this.dialog.open(EditImageComponent, { width: '80%', data });

    // this.dialog.afterAllClosed.subscribe(result => {
    //   this.dataSource.load()
    // })
  }

  openAddImageDialog() {
    const dialogRef = this.dialog.open(AddImageComponent, { width: '80%' });

    dialogRef.afterClosed().subscribe(result => {
      this.dataSource.load()
    });
  }

  syncWithFilesystem() {
    this.apiService.syncImages().subscribe(result => {
      this.dataSource.load()
    })
  }

  ngOnInit(): void {
  }

}