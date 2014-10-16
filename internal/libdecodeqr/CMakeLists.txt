# Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

project(libdecodeqr)

#------------------------------------------------------------------------------
# opencv110
#------------------------------------------------------------------------------

add_definitions(
  -DCVAPI_EXPORTS
)

include_directories(AFTER
	./opencv110/cxcore/include
	./opencv110/cv/include
	./opencv110/cvaux/include
	./opencv110/ml/include
	./opencv110/highgui/include

	./opencv110/cxcore/src
	./opencv110/cv/src
	./opencv110/cvaux/src
	./opencv110/ml/src
	./opencv110/highgui/src
)

set(CXCORE_SRC
	./opencv110/cxcore/src/cxalloc.cpp
	./opencv110/cxcore/src/cxarithm.cpp
	./opencv110/cxcore/src/cxarray.cpp
	./opencv110/cxcore/src/cxcmp.cpp
	./opencv110/cxcore/src/cxconvert.cpp
	./opencv110/cxcore/src/cxcopy.cpp
	./opencv110/cxcore/src/cxdatastructs.cpp
	./opencv110/cxcore/src/cxdrawing.cpp
	./opencv110/cxcore/src/cxdxt.cpp
	./opencv110/cxcore/src/cxerror.cpp
	./opencv110/cxcore/src/cximage.cpp
	./opencv110/cxcore/src/cxjacobieigens.cpp
	./opencv110/cxcore/src/cxlogic.cpp
	./opencv110/cxcore/src/cxlut.cpp
	./opencv110/cxcore/src/cxmathfuncs.cpp
	./opencv110/cxcore/src/cxmatmul.cpp
	./opencv110/cxcore/src/cxmatrix.cpp
	./opencv110/cxcore/src/cxmean.cpp
	./opencv110/cxcore/src/cxmeansdv.cpp
	./opencv110/cxcore/src/cxminmaxloc.cpp
	./opencv110/cxcore/src/cxnorm.cpp
	./opencv110/cxcore/src/cxouttext.cpp
	./opencv110/cxcore/src/cxpersistence.cpp
	./opencv110/cxcore/src/cxprecomp.cpp
	./opencv110/cxcore/src/cxrand.cpp
	./opencv110/cxcore/src/cxsumpixels.cpp
	./opencv110/cxcore/src/cxsvd.cpp
	#./opencv110/cxcore/src/cxswitcher.cpp
	./opencv110/cxcore/src/cxswitcher_new.cpp
	./opencv110/cxcore/src/cxtables.cpp
	./opencv110/cxcore/src/cxutils.cpp
)

set(CV_SRC
	./opencv110/cv/src/cvaccum.cpp
	./opencv110/cv/src/cvadapthresh.cpp
	./opencv110/cv/src/cvapprox.cpp
	./opencv110/cv/src/cvcalccontrasthistogram.cpp
	./opencv110/cv/src/cvcalcimagehomography.cpp
	./opencv110/cv/src/cvcalibinit.cpp
	./opencv110/cv/src/cvcalibration.cpp
	./opencv110/cv/src/cvcamshift.cpp
	./opencv110/cv/src/cvcanny.cpp
	./opencv110/cv/src/cvcolor.cpp
	./opencv110/cv/src/cvcondens.cpp
	./opencv110/cv/src/cvcontours.cpp
	./opencv110/cv/src/cvcontourtree.cpp
	./opencv110/cv/src/cvconvhull.cpp
	./opencv110/cv/src/cvcorner.cpp
	./opencv110/cv/src/cvcornersubpix.cpp
	./opencv110/cv/src/cvderiv.cpp
	./opencv110/cv/src/cvdistransform.cpp
	./opencv110/cv/src/cvdominants.cpp
	./opencv110/cv/src/cvemd.cpp
	./opencv110/cv/src/cvfeatureselect.cpp
	./opencv110/cv/src/cvfilter.cpp
	./opencv110/cv/src/cvfloodfill.cpp
	./opencv110/cv/src/cvfundam.cpp
	./opencv110/cv/src/cvgeometry.cpp
	./opencv110/cv/src/cvhaar.cpp
	./opencv110/cv/src/cvhistogram.cpp
	./opencv110/cv/src/cvhough.cpp
	./opencv110/cv/src/cvimgwarp.cpp
	./opencv110/cv/src/cvinpaint.cpp
	./opencv110/cv/src/cvkalman.cpp
	./opencv110/cv/src/cvkdtree.cpp
	./opencv110/cv/src/cvlinefit.cpp
	./opencv110/cv/src/cvlkpyramid.cpp
	./opencv110/cv/src/cvmatchcontours.cpp
	./opencv110/cv/src/cvmoments.cpp
	./opencv110/cv/src/cvmorph.cpp
	./opencv110/cv/src/cvmotempl.cpp
	./opencv110/cv/src/cvoptflowbm.cpp
	./opencv110/cv/src/cvoptflowhs.cpp
	./opencv110/cv/src/cvoptflowlk.cpp
	./opencv110/cv/src/cvpgh.cpp
	./opencv110/cv/src/cvposit.cpp
	./opencv110/cv/src/cvprecomp.cpp
	./opencv110/cv/src/cvpyramids.cpp
	./opencv110/cv/src/cvpyrsegmentation.cpp
	./opencv110/cv/src/cvrotcalipers.cpp
	./opencv110/cv/src/cvsamplers.cpp
	./opencv110/cv/src/cvsegmentation.cpp
	./opencv110/cv/src/cvshapedescr.cpp
	./opencv110/cv/src/cvsmooth.cpp
	./opencv110/cv/src/cvsnakes.cpp
	./opencv110/cv/src/cvstereobm.cpp
	./opencv110/cv/src/cvstereogc.cpp
	./opencv110/cv/src/cvsubdivision2d.cpp
	./opencv110/cv/src/cvsumpixels.cpp
	./opencv110/cv/src/cvsurf.cpp
	./opencv110/cv/src/cvswitcher.cpp
	./opencv110/cv/src/cvtables.cpp
	./opencv110/cv/src/cvtemplmatch.cpp
	./opencv110/cv/src/cvthresh.cpp
	./opencv110/cv/src/cvundistort.cpp
	./opencv110/cv/src/cvutils.cpp
)

set(CVAUX_SRC
	./opencv110/cvaux/src/camshift.cpp
	./opencv110/cvaux/src/cv3dtracker.cpp
	./opencv110/cvaux/src/cvaux.cpp
	./opencv110/cvaux/src/cvauxutils.cpp
	./opencv110/cvaux/src/cvbgfg_acmmm2003.cpp
	./opencv110/cvaux/src/cvbgfg_codebook.cpp
	./opencv110/cvaux/src/cvbgfg_common.cpp
	./opencv110/cvaux/src/cvbgfg_gaussmix.cpp
	./opencv110/cvaux/src/cvcalibfilter.cpp
	./opencv110/cvaux/src/cvclique.cpp
	./opencv110/cvaux/src/cvcorrespond.cpp
	./opencv110/cvaux/src/cvcorrimages.cpp
	./opencv110/cvaux/src/cvcreatehandmask.cpp
	./opencv110/cvaux/src/cvdpstereo.cpp
	./opencv110/cvaux/src/cveigenobjects.cpp
	./opencv110/cvaux/src/cvepilines.cpp
	./opencv110/cvaux/src/cvface.cpp
	./opencv110/cvaux/src/cvfacedetection.cpp
	./opencv110/cvaux/src/cvfacetemplate.cpp
	./opencv110/cvaux/src/cvfindface.cpp
	./opencv110/cvaux/src/cvfindhandregion.cpp
	./opencv110/cvaux/src/cvhmm.cpp
	./opencv110/cvaux/src/cvhmm1d.cpp
	./opencv110/cvaux/src/cvhmmobs.cpp
	./opencv110/cvaux/src/cvlcm.cpp
	./opencv110/cvaux/src/cvlee.cpp
	./opencv110/cvaux/src/cvlevmar.cpp
	./opencv110/cvaux/src/cvlevmarprojbandle.cpp
	./opencv110/cvaux/src/cvlevmartrif.cpp
	./opencv110/cvaux/src/cvlines.cpp
	./opencv110/cvaux/src/cvlmeds.cpp
	./opencv110/cvaux/src/cvmat.cpp
	./opencv110/cvaux/src/cvmorphcontours.cpp
	./opencv110/cvaux/src/cvmorphing.cpp
	./opencv110/cvaux/src/cvprewarp.cpp
	./opencv110/cvaux/src/cvscanlines.cpp
	./opencv110/cvaux/src/cvsegment.cpp
	./opencv110/cvaux/src/cvsubdiv2.cpp
	./opencv110/cvaux/src/cvtexture.cpp
	./opencv110/cvaux/src/cvtrifocal.cpp
	./opencv110/cvaux/src/cvvecfacetracking.cpp
	./opencv110/cvaux/src/cvvideo.cpp
	./opencv110/cvaux/src/decomppoly.cpp
	./opencv110/cvaux/src/enmin.cpp
	./opencv110/cvaux/src/extendededges.cpp
	./opencv110/cvaux/src/precomp.cpp
)

set(ML_SRC
	./opencv110/ml/src/ml.cpp
	./opencv110/ml/src/ml_inner_functions.cpp
	./opencv110/ml/src/mlann_mlp.cpp
	./opencv110/ml/src/mlboost.cpp
	./opencv110/ml/src/mlcnn.cpp
	./opencv110/ml/src/mlem.cpp
	./opencv110/ml/src/mlestimate.cpp
	./opencv110/ml/src/mlknearest.cpp
	./opencv110/ml/src/mlnbayes.cpp
	./opencv110/ml/src/mlrtrees.cpp
	./opencv110/ml/src/mlsvm.cpp
	./opencv110/ml/src/mltestset.cpp
	./opencv110/ml/src/mltree.cpp
)

set(HIGHGUI_SRC_COMMON
	./opencv110/highgui/src/EasyBMP.cpp
	./opencv110/highgui/src/bitstrm.cpp
	./opencv110/highgui/src/cvcap.cpp
	./opencv110/highgui/src/cvcap_images.cpp
	./opencv110/highgui/src/image.cpp
	./opencv110/highgui/src/jpgd.cpp
	./opencv110/highgui/src/jpge.cpp
	./opencv110/highgui/src/loadsave_new.cpp
	./opencv110/highgui/src/lodepng.cpp
	./opencv110/highgui/src/precomp.cpp
	./opencv110/highgui/src/utils.cpp
	./opencv110/highgui/src/window.cpp
)

if(WIN32 OR MINGW)
	set(HIGHGUI_SRC ${HIGHGUI_SRC_COMMON}
		./opencv110/highgui/src/windows/cvcap_dshow.cpp
		./opencv110/highgui/src/windows/cvcap_vfw.cpp
		./opencv110/highgui/src/windows/window_w32.cpp

		./opencv110/highgui/src/windows/cap_ffmpeg.cpp
		./opencv110/highgui/src/windows/videoInput_capi.cpp
		./opencv110/highgui/src/windows/videoInput.cpp
	)
endif()

if(UNIX AND NOT MINGW)
	set(HIGHGUI_SRC ${HIGHGUI_SRC_COMMON}
		# todo
	)
endif()

if(APPLE)
	set(HIGHGUI_SRC ${HIGHGUI_SRC_COMMON}
		# todo
	)
endif()

if(WIN32 OR MINGW)
	set(OPENCV_LINK_LIBS
		Comctl32 Vfw32 Ole32 Strmiids OleAut32 Uuid Ws2_32
	)
endif()

if(UNIX AND NOT MINGW)
	set(OPENCV_LINK_LIBS
		# todo
	)
endif()

if(APPLE)
	set(OPENCV_LINK_LIBS
		# todo
	)
endif()


set(OPENCV_SRC
	${CXCORE_SRC}
	${CV_SRC}
	${CVAUX_SRC}
	${ML_SRC}
	${HIGHGUI_SRC}
)

#------------------------------------------------------------------------------
# libdecodeqr
#------------------------------------------------------------------------------

include_directories(AFTER
	./libdecodeqr
)

set(libdecodeqr_SRC
	./libdecodeqr/bitstream.cpp
	./libdecodeqr/bitstream.h
	./libdecodeqr/codedata.cpp
	./libdecodeqr/codedata.h
	./libdecodeqr/container.cpp
	./libdecodeqr/container.h
	./libdecodeqr/decodeqr.h
	./libdecodeqr/ecidecoder.cpp
	./libdecodeqr/ecidecoder.h
	./libdecodeqr/formatinfo.cpp
	./libdecodeqr/formatinfo.h
	./libdecodeqr/galois.cpp
	./libdecodeqr/galois.h
	./libdecodeqr/imagereader.cpp
	./libdecodeqr/imagereader.h
	./libdecodeqr/libdecodeqr.cpp
	./libdecodeqr/qrerror.h
	./libdecodeqr/qrtypes.h
	./libdecodeqr/version.h
)

#------------------------------------------------------------------------------

add_library(decodeqrLib SHARED
	${libdecodeqr_SRC}
	${OPENCV_SRC}
)
set_target_properties(decodeqrLib
	PROPERTIES OUTPUT_NAME "decodeqr"
)
target_link_libraries(decodeqrLib
	${OPENCV_LINK_LIBS}
)

install(TARGETS decodeqrLib
	RUNTIME DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
	LIBRARY DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
	ARCHIVE DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
)

#------------------------------------------------------------------------------

add_executable(simple
	./examples/simple/simpletest.cpp
	${libdecodeqr_SRC}
	${OPENCV_SRC}
)
target_link_libraries(simple
	${OPENCV_LINK_LIBS}
)
install(TARGETS simple
	RUNTIME DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
	LIBRARY DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
	ARCHIVE DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
)

add_executable(webcam
	./examples/webcam/webcam.cpp
	${libdecodeqr_SRC}
	${OPENCV_SRC}
)
target_link_libraries(webcam
	${OPENCV_LINK_LIBS}
)
install(TARGETS webcam
	RUNTIME DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
	LIBRARY DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
	ARCHIVE DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
)


#------------------------------------------------------------------------------
