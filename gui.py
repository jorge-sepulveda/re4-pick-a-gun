import sys
import data

from PyQt6.QtWidgets import QApplication, QPushButton, QVBoxLayout, QWidget, QLabel, QHBoxLayout

class Window(QWidget):
    def __init__(self, parent=None):
        super().__init__(parent)
        self.gun = data.Gun()
        self.loadButton = QPushButton(text="Load Game", parent=self)
        self.loadButton.setFixedSize(150, 60)
        self.loadButton.clicked.connect(self.onLoadButtonClicked)

        self.saveButton = QPushButton(text="Save Game", parent=self)
        self.saveButton.setFixedSize(150, 60)
        self.saveButton.clicked.connect(self.onSaveButtonClicked)

        self.rollButton = QPushButton(text="Roll!", parent=self)
        self.rollButton.setFixedSize(150, 60)
        self.rollButton.clicked.connect(self.onrollButtonClicked)

        self.guideButton = QPushButton(text="Guide!", parent=self)
        self.guideButton.setFixedSize(150, 60)
        self.guideButton.clicked.connect(self.onguideButtonClicked)

        # self.quitButton = QPushButton(text="Quit", parent=self)
        # self.quitButton.setFixedSize(150, 60)
        # self.quitButton.clicked.connect(self.onquitButtonClicked)

        self.l = QLabel(text="What are ya buying?")
        self.statusMessage = "Welcome!"
        self.statusMessage = QLabel(text=self.statusMessage)

        verticalLayout = QVBoxLayout()
        horizonalLayout = QHBoxLayout()
        horizonalLayout.addWidget(self.loadButton)
        horizonalLayout.addWidget(self.saveButton)
        horizonalLayout.addWidget(self.rollButton)
        horizonalLayout.addWidget(self.guideButton)
        verticalLayout.addLayout(horizonalLayout)
        verticalLayout.addWidget(self.l)
        verticalLayout.addWidget(self.statusMessage)

        self.setLayout(verticalLayout)

    def onLoadButtonClicked(self):
        if self.gunCheck():
            self.statusMessage.setText(self.gun.load()) 
            self.gun.currentChapter+=1
        else:
            self.statusMessage.setText("Load What Stranger, you're done!")

    def onSaveButtonClicked(self):
        if self.gunCheck():
            self.statusMessage.setText(self.gun.save()) 
        else:
            self.statusMessage.setText("No Saving here, Stranger. Go Home.")

    def onrollButtonClicked(self):
        if self.gunCheck():
            self.statusMessage.setText(self.gun.roll()) 
            self.gun.currentChapter+=1
        else:
            self.statusMessage.setText("Out of guns, Stranger.")

    def onguideButtonClicked(self):
        self.gun.guide()

    def gunCheck(self):
        return self.gun.currentChapter < self.gun.finalChapter

if __name__ == "__main__":
    app = QApplication(sys.argv)
    window = Window()
    window.show()
    sys.exit(app.exec())
